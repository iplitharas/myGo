package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	_, err := fmt.Fprintf(writer, "Hello, "+name)
	if err != nil {
		return
	}

}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "dependency injection")
	//log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
