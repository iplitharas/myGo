package main

import (
	"log"
	"net/http"
	"web-app/pkg/handlers"
)

func main() {
	err := http.ListenAndServe("localhost:8080", handlers.CreateMux())
	if err != nil {
		log.Fatalf("error trying to start the server: %s", err)
	}
}
