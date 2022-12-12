package main

import (
	"errors"
	"fmt"
)

type StatusError struct {
	Status  int
	Message string
	err     error
}

// StatusError implementation for the Error interface
func (se StatusError) Error() string {
	return se.Message
}

// Unwrap StatusError implementation for the Unwrap interface
func (se StatusError) Unwrap() error {
	fmt.Println("------Calling Unwrap method-----")
	return se.err
}

func main() {
	fmt.Println("Wrapping error with custom error types")
	// error without any wrapped err
	customError := StatusError{
		Status:  0,
		Message: "Some error here",
		err:     nil,
	}
	fmt.Printf("Custom error is: %s\n", customError)
	if wrappedErr := errors.Unwrap(customError); wrappedErr != nil {
		fmt.Printf("Wrapped Error is: %s\n", wrappedErr)
	} else {
		fmt.Println("There isn't any wrapped error..")
	}

	// error with an wrapped error
	customError = StatusError{
		Status:  0,
		Message: "Some error here",
		err:     fmt.Errorf("wrapped error here"),
	}
	fmt.Printf("Custom error is: %s\n", customError)
	if wrappedErr := errors.Unwrap(customError); wrappedErr != nil {
		fmt.Printf("Wrapped Error is: %s\n", wrappedErr)
	} else {
		fmt.Println("There isn't any wrapped error..")
	}

}
