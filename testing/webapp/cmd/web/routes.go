package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	//register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIPtoContext)

	// register routes
	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)

	fileserver := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))
	// static assets

	return mux
}
