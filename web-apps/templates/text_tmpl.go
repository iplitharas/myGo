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
	err := tmp.ExecuteTemplate(os.Stdout, "tmpl.html", 123)
	if err != nil {
		fmt.Println("error is:", err)
	}
}
