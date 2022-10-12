package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/render"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig

// session is declared as package level variable
// to be  used from middleware,
// for the handlers we added the session in the config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"

// main is the main application function
func main() {
	app.InProduction = false
	// create a big session
	// it stores the session  by default at cookies
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// to use https or http !
	session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()
	app.TemplateCache = tc
	app.UseCache = false
	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	srv := &http.Server{
		Addr:              portNumber,
		Handler:           routes(&app),
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	//err = http.ListenAndServe(portNumber, nil)
	//if err != nil {
	//	return
	//}
}
