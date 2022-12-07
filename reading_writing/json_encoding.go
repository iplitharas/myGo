package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{
		Name: "Ioannis",
		Age:  32,
	}
	fmt.Println(person)
	tmpFile, err := os.CreateTemp(os.TempDir(), "json")
	defer tmpFile.Close()
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(tmpFile).Encode(person)
	if err != nil {
		panic(err)
	}
	// Reset the pointer to the start of the file
	tmpFile.Seek(0, 0)
	var fromFile Person
	err = json.NewDecoder(tmpFile).Decode(&fromFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("person back")
	fmt.Println(fromFile)
}
