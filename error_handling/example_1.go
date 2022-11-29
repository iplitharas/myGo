package main

import (
	"errors"
	"fmt"
)

// calcRemainderAndMod divides the two numbers and returns
// the remainder and the modulo.
// in uses the `errors.New` to create the error
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot perform division with 0")
	}
	return numerator / denominator, numerator & denominator, nil
}

func main() {
	fmt.Println("Errors in Golang")
	numerator, denominator, err := calcRemainderAndMod(10, 0)
	if err != nil {
		fmt.Print("Error is: ", err)
	} else {
		fmt.Printf("10/2= %d  \n10 mod 2= %d\n", numerator, denominator)
	}

}
