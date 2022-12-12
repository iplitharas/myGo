package main

import "fmt"

type Sentinel string

// Implement the Error interface
// for the custom type `Sentinel`
func (s Sentinel) Error() string {
	return string(s)
}

const ErrFoo = Sentinel("foo error")
const ErrBar = Sentinel("bar error")

func returnSentinelErrors() (error, error) {
	fmt.Println("Hi, there, here are two errors for you!")
	return ErrFoo, ErrBar
}

func main() {
	fmt.Println("Sentinel Errors in Go!")
	firstErr, secondErr := returnSentinelErrors()
	fmt.Printf("First error is: %s\nSecond error is: %s\n", firstErr, secondErr)

}
