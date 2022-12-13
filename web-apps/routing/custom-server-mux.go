package main

import (
	"fmt"
	"log"
	"net/http"
)

type indexHandler struct{}

func (indexHandler indexHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "index page")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello there")
	})
	mux.Handle("/index", indexHandler{})
	srv := http.Server{Addr: ":8080", Handler: mux}
	log.Printf("Listening at: http://localhost:8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
