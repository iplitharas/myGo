package main

import (
	"log"
	"net/http"
	"text/template"
)

type indexHandler struct{}

func (index indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("error during parsing the form")
	}
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("error during parsing the template: %q", err)
	}
	err = tmp.ExecuteTemplate(w, "index.html", r.PostForm)
	if err != nil {
		log.Fatalf("error during exetute of template: %q", err)
	}
}

func main() {

	srv := http.Server{Addr: ":8081", Handler: indexHandler{}}
	log.Println("Listening at: http://localhost:8081")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}
}
