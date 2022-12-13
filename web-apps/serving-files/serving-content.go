package main

import (
	"log"
	"net/http"
	"os"
)

func photoCopy(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./static/london.png")
	if err != nil {
		log.Printf("cannot open the image: %q", err)
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}
	log.Printf("Listening at: http://localhost:8080")
	http.HandleFunc("/", photoCopy)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
