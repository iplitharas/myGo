package main

import (
	"fmt"
	"net/http"
	"web-programming/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
