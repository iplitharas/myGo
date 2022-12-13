package main

import (
	"io"
	"log"
	"net/http"
)

func photoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="london.png">`)
}
func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}
	log.Printf("Listening at: http://localhost:8080")
	http.HandleFunc("/", photoHandler)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
