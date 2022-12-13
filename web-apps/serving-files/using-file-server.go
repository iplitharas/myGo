package main

import (
	"io"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/static/london.png">`)
}

func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}

	log.Printf("Listening at: http://localhost:8080")
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/london/", indexHandler)
	http.Handle("/favicon.ico/", http.NotFoundHandler())
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
