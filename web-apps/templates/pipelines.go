package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}
func addMessage(str string) string {
	return str + " Go go go"
}

var functions = template.FuncMap{"fdayMDY": monthDayYear, "addMessage": addMessage}

func main() {
	// Parse the templates
	tmp := template.Must(template.New("").Funcs(functions).ParseFiles("templates/pipelines.html"))

	// write the template to `os.Stdout`

	err := tmp.ExecuteTemplate(os.Stdout, "pipelines.html", time.Now())
	if err != nil {
		fmt.Println("error is:", err)
	}
}
