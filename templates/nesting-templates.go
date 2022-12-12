package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Parse the templates
	tmp := template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))

	// write the template to `os.Stdout`

	err := tmp.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		fmt.Println("error is:", err)
	}
}
