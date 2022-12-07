package main

import "fmt"

func main() {
	var firstName, lastName string
	fmt.Println("Enter your first and lastname")
	_, err := fmt.Scanln(&firstName, &lastName)
	if err != nil {
		fmt.Println("Error trying to read the input from the keyboard")
	}
	fmt.Printf("Hello %q %q\n", firstName, lastName)
}
