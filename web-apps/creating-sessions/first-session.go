package main

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"text/template"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // UserID, user
var dbSessions = map[string]string{} //SessionID, User ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	// get Cookie or create a new one
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Printf("session-id is missing: %q", err)
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}
		log.Printf("Creating new session: %q", id)
		http.SetCookie(w, cookie)
	}
	// if the user already exists, get user
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}
	err = tpl.ExecuteTemplate(w, "index.html", u)
	if err != nil {
		log.Printf("error during executing the template %q", err)
	}
}
func bar(w http.ResponseWriter, r *http.Request) {
	// get cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	user := dbUsers[un]
	err = tpl.ExecuteTemplate(w, "bar.html", user)
	if err != nil {
		log.Printf("error during executing the template %q", err)
	}

}

func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}
	log.Printf("Listening at: http://localhost:8080")
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}
}
