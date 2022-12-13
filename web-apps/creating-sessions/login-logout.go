package main

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"text/template"
	"time"
)

type userID string

type UserInfo struct {
	userID    userID
	FirstName string
	LastName  string
}

var dbCookiesStore = make(map[string]userID)
var dbUsersStore = make(map[userID]UserInfo)

func login(w http.ResponseWriter, r *http.Request) {

	if isAlreadyLogin(r) {
		// nothing to do
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.Method == http.MethodGet {
		tmp, err := template.ParseGlob("./templates-login/**")
		if err != nil {
			log.Fatalf("error during parsing the templates dir: %q", err)
		}
		tmp.ExecuteTemplate(w, "login.html", nil)
		return
	}
	// render the form and create a new cookie
	firstName := r.FormValue("fname")
	lastName := r.FormValue("lname")
	log.Printf("first name  is: %q and last name is:%q", firstName, lastName)
	user := UserInfo{
		userID:    userID(uuid.NewV4().String()),
		FirstName: firstName,
		LastName:  lastName,
	}
	// store the user to the database
	dbUsersStore[user.userID] = user
	// create a session for this user
	id := uuid.NewV4()
	cookie := &http.Cookie{
		Name:     "session",
		Value:    id.String(),
		HttpOnly: true,
	}
	// store the user id to the session - user id mapping
	dbCookiesStore[id.String()] = user.userID
	log.Printf("Creating new session: %q", id)
	http.SetCookie(w, cookie)
	// Redirect user to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func logout(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("./templates-login/**")
	if err != nil {
		log.Fatalf("error during parsing the templates dir: %q", err)
	}
	if !isAlreadyLogin(r) {
		tmp.ExecuteTemplate(w, "home.html", nil)
		return
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	}

	log.Printf("Removing the session!")
	http.SetCookie(w, cookie)
	// Redirect user to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func home(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("./templates-login/**")
	if err != nil {
		log.Fatalf("error during parsing the templates dir: %q", err)
	}

	if !isAlreadyLogin(r) {
		tmp.ExecuteTemplate(w, "home.html", nil)
		return
	}
	// fetch the user
	cookie, err := r.Cookie("session")
	userId := dbCookiesStore[cookie.Value]
	user := dbUsersStore[userId]
	tmp.ExecuteTemplate(w, "home.html", user)

}

func main() {
	srv := http.Server{Addr: ":8080", Handler: nil}
	log.Printf("Listening at: http://localhost:8080")
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/", home)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to server the server: %q", err)
	}
}

func isAlreadyLogin(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Printf("Trying to access: %q without cookie: %q", r.Method, err)
		return false
	}
	if _, ok := dbCookiesStore[cookie.Value]; !ok {
		log.Printf("Cookie: %q is not associted with any user id", cookie.Value)
		return false
	}
	if userID, ok := dbCookiesStore[cookie.Value]; ok {
		log.Printf("Cookie: %q is associted with user id: %q", cookie.Value, userID)
	}
	return true

}
