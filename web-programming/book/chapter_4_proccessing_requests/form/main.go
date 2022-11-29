package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering formHandler")
	err := r.ParseForm()
	if err != nil {
		log.Println("error during parsing the form", err)
		return
	}
	fmt.Fprintln(w, r.Form)
}

func main() {
	log.Println("Form data demo")
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/form", formHandler)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("error tryign to serve the server! %s", err)
		return
	}
}
