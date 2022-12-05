package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var testCases = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound},
	}
	var app application
	routes := app.routes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()
	pathToTemplates = "./../../templates"
	for _, testCase := range testCases {
		response, err := ts.Client().Get(ts.URL + testCase.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if response.StatusCode != testCase.expectedStatusCode {
			t.Errorf("for test case: %s\nexpected %d but received %d",
				testCase.name, testCase.expectedStatusCode, response.StatusCode)
		}
	}
}
