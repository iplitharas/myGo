package main

import (
	"fmt"
	"log"
	"net/http"
	"web-programming/pkg/config"
	"web-programming/pkg/handlers"
	"web-programming/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
