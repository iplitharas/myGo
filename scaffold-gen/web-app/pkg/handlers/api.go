package handlers

import (
	"fmt"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached Echorand Corp’s Service API” as the response")
}
