package main

import (
	"fmt"
	"log"
	"net/http"
)

type ServerHandler struct{}

func (sh ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/index":
		fmt.Fprintf(w, "Index page")
	case "/hello":
		fmt.Fprintf(w, "Hello there")
	default:
		fmt.Fprintf(w, "I'm a custom mux")

	}
}

func main() {
	srv := http.Server{Addr: ":8080", Handler: ServerHandler{}}
	log.Printf("Listening at: http://localhost:8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
