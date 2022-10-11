## Basic web application in Go

```go
package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello world!")
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

```

## Render templates
```go
package main

import (
"fmt"
"html/template"
"net/http"
)
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// it checks always from the root
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

```