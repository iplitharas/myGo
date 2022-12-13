package main

import (
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func createUUID(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Printf("session-id is missing: %q", err)
		id := uuid.NewV4()
		cookie := &http.Cookie{Name: "session",
			Value: id.String(),
			// we can't access this cookie from javascript for example
			HttpOnly: true}
		log.Printf("Creating new session: %q", id)
		http.SetCookie(w, cookie)
	}
	log.Printf("Using existing session: %q", cookie)
}

func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}
	log.Printf("Listening at: http://localhost:8080")
	http.HandleFunc("/create", createUUID)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
