package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil

}

func main() {
	fmt.Println("Checking error with `error.Is`")
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Printf("Error is: %s\n", err)
	}
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("That file doesn't exist\n")
	}
	if wrappedError := errors.Unwrap(err); wrappedError != nil {
		fmt.Printf("Wrapped error is: %s\n", wrappedError)
	}
}
