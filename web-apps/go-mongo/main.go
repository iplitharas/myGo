package main

import (
	json2 "encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-mongo/models"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello")

}

// getUser Handler
func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	user := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    "32",
		Id:     p.ByName("id"),
	}
	err := json2.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		errorMessage := fmt.Sprintf("error trying to serialize user with id:"+
			" %q is: %q", p.ByName("id"), err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
	}
}

func main() {

	router := httprouter.New()
	srv := http.Server{Addr: ":8080", Handler: router}
	log.Println("Listening at: http://localhost:8080")
	router.GET("/", index)
	router.GET("/user/:id", getUser)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to start the server: %q", err)
	}
}
