package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

func writeCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,
		&http.Cookie{Name: "my-cookie", Value: "chocolate"})
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func deleteCookie(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:    "my-cookie",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}
func index(w http.ResponseWriter, r *http.Request) {
	tmp := template.New("home")
	cookies := r.Cookies()
	htmlContent := `
<html>
<h1> Cookies demo</h1>
<h3><a href="http://localhost:8080/write">Write a cookie</a> <h3>
<h3> <a href="http://localhost:8080/delete">Delete a cookie</a><h3>
 {{ if . }}
 	{{ range  . }}
		<ul> 
		<li> {{ .Name}} {{ .Value }} </li> 
		</ul>
		{{ end }} 
{{else}} 
	<h2> No cookies </h2> 
{{ end }} 
</html>`
	_, err := tmp.Parse(htmlContent)
	if err != nil {
		log.Printf("error during parsing the template: %q", err)

	}
	err = tmp.Execute(w, cookies)
	if err != nil {
		log.Printf("error during executing the template: %q", err)
	}
	//io.WriteString(w, htmlContent)
}

func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}
	log.Printf("Listening at: http://localhost:8080")
	http.HandleFunc("/", index)
	http.HandleFunc("/write", writeCookie)
	http.HandleFunc("/delete", deleteCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}

}
