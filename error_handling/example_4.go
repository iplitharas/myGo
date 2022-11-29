package main

import "fmt"

// Since error is an interface, you can define your own errors that include additional information
// for logging or error handling. For example you might want to include a status code as part of the error to
// indicate the kind of error that should be reported back to the user.
// This lets you avoid string comparisons to determine error causes.

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
	NotExist
)

// CustomError struct to hold
// the status code
// and the actual error message
type CustomError struct {
	Status  Status
	Message string
}

// implement the Error interface
func (ce CustomError) Error() string {
	return fmt.Sprintf("Error is: %s with status code: %d", ce.Message, ce.Status)
}

// returnCustomErrors returns our error, but it's not good practice
// we should always use the error type instead of the CustomError
// in that case we'll lose the information of the status
func returnCustomErrors() (error, error, error) {
	fmt.Println("Hello, here are three custom errors for you!")
	invalidLoginErr := CustomError{
		Status:  InvalidLogin,
		Message: "Invalid login ",
	}
	notFoundErr := CustomError{
		Status:  NotFound,
		Message: "Not found",
	}
	notExistErr := CustomError{
		Status:  NotExist,
		Message: "Not exist",
	}
	return invalidLoginErr, notFoundErr, notExistErr
}

func main() {
	fmt.Println("Custom errors in Go!")
	firstErr, _, _ := returnCustomErrors()
	fmt.Printf("First error is: %s\n", firstErr.Error())

}
