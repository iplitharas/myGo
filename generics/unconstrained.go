package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	id        int
}

func addUser[T any](users []T, user T) []T {
	return append(users, user)

}

func main() {
	// add some str
	usersString := make([]string, 0)
	usersString = addUser(usersString, "Ioannis")
	usersString = addUser(usersString, "Michael")
	fmt.Println("Users (strings) are: ", usersString)
	// add some int
	userInts := make([]int, 0)
	userInts = addUser(userInts, 1)
	userInts = addUser(userInts, 2)
	fmt.Println("Users (ints) are: ", userInts)
	users := make([]User, 0)
	users = addUser(users, User{
		firstName: "Ioannis",
		lastName:  "Plitharas",
		id:        1,
	})
	fmt.Println("Users  are: ", users)

}
