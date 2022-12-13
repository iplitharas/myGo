package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func London(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="london.png">`)
}

func LondonCopy(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./static/london.png")
	if err != nil {
		log.Printf("cannot open the image: %q", err)
	}
	defer f.Close()
	io.Copy(w, f)
}
func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}
	log.Printf("Listening at: http://localhost:8080")
	http.HandleFunc("/", London)
	http.HandleFunc("/copy", LondonCopy)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
