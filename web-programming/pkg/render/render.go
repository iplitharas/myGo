package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Println("not able to find template within a cache")
	}

	buff := new(bytes.Buffer)
	err = t.Execute(buff, nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("using template cache for: %s", tmpl)
	// render the template
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all the files named *page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		log.Println(err)
		return myCache, err
	}
	// range through all files
	for _, page := range pages {
		fileName := filepath.Base(page)
		ts, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
		}
		myCache[fileName] = ts
	}
	return myCache, nil

}
