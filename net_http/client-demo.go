package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	// create a simple client
	client := &http.Client{
		Transport: nil,
		Timeout:   30 * time.Second,
	}

	// Create a request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-My-Client", "Learning GO")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println(res)

}
