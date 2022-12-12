package main

import (
	"fmt"
	"log"
	"net/http"
)

type index struct{}

func (in index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello there")
}

func main() {
	index := index{}
	srv := http.Server{Addr: ":8081", Handler: index}
	log.Println("Listening at: http://localhost:8081")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}

}
