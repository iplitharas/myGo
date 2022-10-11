package main

import "fmt"

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	myFunc(MyFuncOpts{
		FirstName: "1",
		LastName:  "2",
		Age:       3,
	})
}

func myFunc(opts MyFuncOpts) error {
	fmt.Println("Called myFunc with:", opts)

	return nil
}
