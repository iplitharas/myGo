package main

import (
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src=https://www.freecodecamp.org/news/content/images/2021/10/golang.png>`)
}

func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}
	log.Printf("Listening at: http://localhost:8080")
	http.HandleFunc("/", index)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
