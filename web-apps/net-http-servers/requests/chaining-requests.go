package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	// the browsers will call the / + the /favicon.ico
	// that's why we are seeing two calls of this
	// handler
	r.ParseForm()
	log.Printf("Entering mainPage from %s\n", r.URL)
	fmt.Fprintf(w, "Hello world\n")
	log.Printf("Leaving mainPage from %s\n", r.URL)

}

func timeRequest(h http.HandlerFunc) http.HandlerFunc {

	name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		log.Printf("timeRequest will call %s", name)
		h(w, r)
		fmt.Fprintf(w, "Request  %q\n", time.Now().Sub(now))
	}
}

func main() {

	srv := http.Server{Addr: ":8081"}
	log.Println("Listening at: http://localhost:8081")
	http.HandleFunc("/", timeRequest(mainPageHandler))
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to start the server: %q", err)
	}
}
