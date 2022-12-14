package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	log.Println(email, password)
	fmt.Fprintf(w, email)
}

var pathToTemplates = "./templates/"

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, tmp string, data *TemplateData) error {
	// parse the template
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, tmp))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}
	data.IP = app.ipFromContext(r.Context())
	// execute the template
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}

	return nil

}
