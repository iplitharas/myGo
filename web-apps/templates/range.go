package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Parse the templates
	tmp := template.Must(template.ParseGlob("templates/*"))

	// write the template to `os.Stdout`

	// slice of strings
	countries := []string{"France", "UK", "GREECE"}
	err := tmp.ExecuteTemplate(os.Stdout, "range.html", countries)
	if err != nil {
		fmt.Println("error is:", err)
	}
}
