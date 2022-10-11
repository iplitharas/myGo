package main

import (
	"fmt"
)

type Person struct {
	name    string
	address string
	age     int
}

func main() {
	p1 := Person{name: "Ioannis", address: "E162UN", age: 32}
	fmt.Println(p1)

	// new always create a pointer
	p2 := new(Person)
	p2.name = "ioannis"
	p2.address = "E16 2UN"
	p2.age = 32
	fmt.Println(p2)

	var p3 Person
	p3.name = "ioannis"
	p3.address = "E16 2UN"
	p3.age = 32
	fmt.Println(p3)

	persons := make([]Person, 3)
	for i := 0; i < 3; i++ {
		persons[i] = Person{name: "ioannis",
			address: "E16 2UN",
			age:     33 + i}

	}
	for idx, person := range persons {
		fmt.Println(idx, person)
	}

}
