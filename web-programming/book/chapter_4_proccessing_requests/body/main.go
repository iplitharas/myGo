package main

import (
	"fmt"
	"log"
	"net/http"
)

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Reading the body of the request")
	bodyLength := r.ContentLength
	log.Printf("Content length is: %d", bodyLength)
	body := make([]byte, bodyLength)
	_, err := r.Body.Read(body)
	if err != nil {
		log.Printf("error during read", err)
	}
	_, err = fmt.Fprintf(w, string(body))
	if err != nil {
		log.Printf("error during writing", err)
	}
}

func main() {
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/body", bodyHandler)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
