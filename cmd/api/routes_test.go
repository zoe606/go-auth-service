package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func TestRoutesExist(t *testing.T) {
	testApp := Config{}

	testRoute := testApp.routes()
	chiRoutes := testRoute.(chi.Router)

	routes := []string{"/authenticate"}
	for _, route := range routes {
		routeExist(t, chiRoutes, route)
	}
}

func routeExist(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}

		return nil
	})

	if !found {
		t.Errorf("didn't find %s in registered routes", route)
	}
}
