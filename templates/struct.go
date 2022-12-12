package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Parse the templates
	tmp := template.Must(template.ParseFiles("templates/struct.html"))

	// write the template to `os.Stdout`

	// slice of strings
	countries := []struct {
		Country string
		Capital string
	}{{"UK", "London"}, {"France", "Paris"}, {"Greece", "Athens"}}
	err := tmp.ExecuteTemplate(os.Stdout, "struct.html", countries)
	if err != nil {
		fmt.Println("error is:", err)
	}
}
