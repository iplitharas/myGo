package main

import (
	json2 "encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

var persons = make([]Person, 0, 3)

func personsMarshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json, err := json2.Marshal(persons)
	if err != nil {
		log.Fatalf("error during encoding the response using marshall: %q", err)
	}
	w.Write(json)
}

func personUnMarshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := `[{"FirstName":"James","LastName":"Bond","Age":32},{"FirstName":"Encoding","LastName":"Unmarshal","Age":42}]`
	err := json2.Unmarshal([]byte(data), &persons)
	if err != nil {
		log.Fatalf("error during decoding the request using marshall: %q", err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func personsEncoding(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json2.NewEncoder(w).Encode(persons)
	if err != nil {
		log.Fatalf("error during encoding the response using json.encoder:%q", err)
	}
}

func personDecoding(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := `[{"FirstName":"James","LastName":"Bond","Age":10},{"FirstName":"Encoding","LastName":"Json","Age":42}]`
	err := json2.NewDecoder(strings.NewReader(data)).Decode(&persons)
	if err != nil {
		log.Fatalf("error during decoding the request using json.encoder:%q", err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func index(w http.ResponseWriter, r *http.Request) {
	homePage := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Go Json encoding decoding</title>
</head>
<body>
<h2> GO- Json Encoding/Decoding</h2>
<ul>
<li><a href="/persons-post-unmarshal">POST persons and decoding using Unmarshal</a></li>
<li><a href="/persons-post-decode">POST persons and decoding using json.Decode</a></li>
<li><a href="/persons-get-json">Get persons using marshal</a></li>
<li><a href="/persons-get-encoding">Get persons using json.Encoder</a></li>
<ul>
</body>
</html>
`
	fmt.Fprintf(w, homePage)
}

func main() {
	srv := http.Server{Addr: ":8081", Handler: nil}
	http.HandleFunc("/", index)
	http.HandleFunc("/persons-post-unmarshal", personUnMarshal)
	http.HandleFunc("/persons-post-decode", personDecoding)
	http.HandleFunc("/persons-get-json", personsMarshal)
	http.HandleFunc("/persons-get-encoding", personsEncoding)
	log.Println("Listening at: http://localhost:8081")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}

}
