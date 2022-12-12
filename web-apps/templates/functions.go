package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(str string) string {
	s := strings.TrimSpace(str)
	s = str[:3]
	return s
}
func main() {
	// Parse the templates
	tmp := template.Must(template.New("").Funcs(fm).ParseFiles("templates/functions.html"))

	// write the template to `os.Stdout`

	// slice of strings
	countries := make(map[string]string)
	countries["Greece"] = "Athens"
	countries["France"] = "Paris"
	countries["UK"] = "London"
	err := tmp.ExecuteTemplate(os.Stdout, "functions.html", countries)
	if err != nil {
		fmt.Println("error is:", err)
	}
}
