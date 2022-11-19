package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello from handler function!")
	if err != nil {
		return
	}
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	log.Printf("Server is running at: %s", server.Addr)
	// register the handler
	http.HandleFunc("/hello", hello)
	err := server.ListenAndServe()
	if err != nil {
		log.Panicln(err)
		return
	}
}
