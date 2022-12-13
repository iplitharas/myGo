package main

import (
	"log"
	"net/http"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call bar using %s", r.URL.Path)

}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call foo using %s", r.URL.Path)
	http.Redirect(w, r, "/bar", http.StatusPermanentRedirect)
}

func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}
	log.Printf("Listening at: http://localhost:8080")
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/foo", fooHandler)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
