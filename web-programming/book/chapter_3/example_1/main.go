package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello World", r.RequestURI)
	if err != nil {
		return
	}
}
func main() {
	handler := myHandler{}
	server := http.Server{
		Addr:              "127.0.0.1:8000",
		Handler:           &handler,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Printf("Server is running")
}
