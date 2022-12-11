package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Parse the template
	tmp := template.Must(template.ParseGlob("templates/*"))

	// write the template to `os.Stdout`
	err := tmp.ExecuteTemplate(os.Stdout, "tmpl.html", nil)
	if err != nil {
		fmt.Println("error is:", err)
	}
}
