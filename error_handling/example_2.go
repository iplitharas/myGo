package main

import (
	"fmt"
)

// calcRemainderAndModV2 divides the two numbers and returns
// the remainder and the modulo.
// in uses the `fmt.ErrorF` to create the error
func calcRemainderAndModV2(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, fmt.Errorf("cannot perform division for numerator: %d and denominator: %d",
			numerator, denominator)
	}
	return numerator / denominator, numerator & denominator, nil
}

func main() {
	fmt.Println("Errors in Golang")
	numerator, denominator, err := calcRemainderAndModV2(10, 0)
	if err != nil {
		fmt.Print("Error is: ", err)
	} else {
		fmt.Printf("10/2= %d  \n10 mod 2= %d\n", numerator, denominator)
	}

}
