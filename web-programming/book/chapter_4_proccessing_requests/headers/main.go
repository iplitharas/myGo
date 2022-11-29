package main

import (
	"log"
	"net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	log.Println("Headers are:", headers)
	h := r.Header["Accept-Encoding"]
	log.Println("Accept-Encoding is:", h)

}

func main() {
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/", headersHandler)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
