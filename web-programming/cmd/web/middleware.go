package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

// WriteToConsole it's a toy middleware to log a message for each request
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running middleware here")
		fmt.Println("Continue with the next middleware")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/", Secure: false,
		SameSite: http.SameSiteLaxMode})
	return csrfHandler
}
