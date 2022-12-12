package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Parse the templates
	tmp := template.Must(template.ParseFiles("templates/map.html"))

	// write the template to `os.Stdout`

	// slice of strings
	countries := make(map[string]string)
	countries["Greece"] = "Athens"
	countries["France"] = "Paris"
	countries["UK"] = "London"
	err := tmp.ExecuteTemplate(os.Stdout, "map.html", countries)
	if err != nil {
		fmt.Println("error is:", err)
	}
}
