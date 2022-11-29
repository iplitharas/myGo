package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

// Chaining Handlers

// mainPageHandler simple handlerFunction
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func logRequest(h http.HandlerFunc) http.HandlerFunc {
	log.Println("Entering Log logRequest")
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		log.Println("Handler function called:", name)
		h(w, r)
	}

}

func LogHeaders(h http.HandlerFunc) http.HandlerFunc {
	log.Println("Entering Log headers")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Referer header is:", r.Header.Get("Referer"))
		h(w, r)
	}
}

func main() {
	log.Println("Chaining handlers example!")
	log.Println("Serving the server at: http://localhost:8080")
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/", LogHeaders(logRequest(mainPageHandler)))
	err := server.ListenAndServe()
	if err != nil {
		log.Panicln("error trying to start the server is:", err)
		return
	}

}
