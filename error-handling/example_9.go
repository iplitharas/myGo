package main

import (
	"errors"
	"fmt"
)

func doThing1() error {
	return fmt.Errorf("an error from doThing1")
}

func doThing2() error {
	return fmt.Errorf("an error from doThing2")
}

func doThing3() error {
	return fmt.Errorf("an error from doThing3")
}

func doSomeThings() (string, error) {
	err1 := doThing1()
	if err1 != nil {
		return "", fmt.Errorf("in doSomeThings: %w", err1)
	}
	err2 := doThing2()
	if err2 != nil {
		return "", fmt.Errorf("in doSomeThings: %w", err2)
	}
	err3 := doThing3()
	if err2 != nil {
		return "", fmt.Errorf("in doSomeThings: %w", err3)
	}

	return "No errors!", nil
}

func doSomethingDefer() (_ string, err error) {
	defer func() {
		if err != nil {
			// we have just to assign the error to the
			// `` err
			fmt.Println("Running defer now!")
			err = fmt.Errorf("in doSomethingDefer: %w", err)
		}
	}()
	err = doThing1()
	if err != nil {
		fmt.Println("Error from doThing1()")
		return "", err
	}
	return "No error", nil
}

func main() {
	fmt.Println("Wrapping errors with defer!")
	_, err := doSomeThings()
	if err != nil {
		fmt.Println("Error without using defer is: ", err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Printf("Wrapped Error is: %s\n", wrappedErr)
		}
	}
	_, err2 := doSomethingDefer()
	if err2 != nil {
		fmt.Println("Error using the defer is:", err2)
		if wrappedErr := errors.Unwrap(err2); wrappedErr != nil {
			fmt.Printf("Wrapped Error is: %s\n", wrappedErr)
		}
	}

}
