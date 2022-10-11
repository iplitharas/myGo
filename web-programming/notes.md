## Basic web application in Go

```go
package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello world!")
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

```

or
```go
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
```

## Render templates
```go
package main

import (
"fmt"
"html/template"
"net/http"
)
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// it checks always from the root
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

```

## Routing
```go
package main

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux

}
```

## Middlewares
```go
// WriteToConsole it's a toy middleware to log a message for each request
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running middleware here")
		fmt.Println("Continue with the next middleware")
		next.ServeHTTP(w, r)
	})
}

```