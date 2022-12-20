package main

import (
	"database/sql"
	// we have to import it because it will be register to sql
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/users?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to postgres!")
}
