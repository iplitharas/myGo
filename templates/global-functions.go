package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Parse the templates
	tmp := template.Must(template.ParseFiles("templates/global-functions.html"))

	xs := []string{"zero", "one", "two", "three"}
	// write the template to `os.Stdout`

	err := tmp.ExecuteTemplate(os.Stdout, "global-functions.html", xs)
	if err != nil {
		fmt.Println("error is:", err)
	}
}
