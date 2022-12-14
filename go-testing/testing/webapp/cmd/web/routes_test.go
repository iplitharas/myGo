package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"testing"
)

func Test_application_routes(t *testing.T) {
	registered := []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/static/*", "GET"},
	}

	var app application
	mux := app.routes()
	chiRoutes := mux.(chi.Routes)
	for _, route := range registered {
		// check if the route exists
		if !routeExist(route.route, route.method, chiRoutes) {
			t.Errorf("route %s is not registered", route.route)
		}
	}

}

func routeExist(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false
	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(handler2 http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})
	return found

}
