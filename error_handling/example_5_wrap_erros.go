package main

import (
	"errors"
	"fmt"
	"os"
)

func wrapError(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in `wrapError` function while trying to open the file: %w", err)
	}
	err = f.Close()
	if err != nil {
		return fmt.Errorf("in `wrapError` function while trying to close the file: %w ", err)
	}
	return nil
}

func main() {
	fmt.Println("Wrapping errors in Go")
	err := wrapError("Not_here.txt")
	if err != nil {
		fmt.Printf("Error is: %s\n", err)
		// check for any wrapped error
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Printf("Wrapped Error is: %s\n", wrappedErr)
		}
	}
}
