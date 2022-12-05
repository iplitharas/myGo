package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIPtoContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}
	var app application
	// create a dummy handler that we'll use to check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure that the value exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}
		// make sure we got a string back
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		t.Log(ip)
	})
	for _, test := range tests {
		//create the handler to test
		handlerToTest := app.addIPtoContext(nextHandler)
		request := httptest.NewRequest("GET", "http:://testing", nil)
		if test.emptyAddr {
			request.RemoteAddr = ""
		}
		if len(test.headerName) > 0 {
			request.Header.Add(test.headerName, test.headerValue)
		}
		if len(test.addr) > 0 {
			request.RemoteAddr = test.addr
		}
		handlerToTest.ServeHTTP(httptest.NewRecorder(), request)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	// create the app
	var app application
	// get a context
	var ctx = context.Background()
	ctx = context.WithValue(ctx, contextUserKey, "test_ip")
	ip := app.ipFromContext(ctx)
	if !strings.EqualFold("test_ip", ip) {
		t.Errorf("expected ip from ipFromContext but got: %s", ip)
	}

}
