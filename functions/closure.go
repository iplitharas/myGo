package main

import (
	"fmt"
	"math"
)

func raiseToPower(power float64) func(number float64) float64 {

	return func(number float64) float64 {
		return math.Pow(number, power)
	}
}

func Apply(numbers []float64, myFunc func(float64) float64) []float64 {
	result := make([]float64, 0, len(numbers))
	for _, number := range numbers {
		result = append(result, myFunc(number))
	}

	defer func() {
		fmt.Println("Running defer block")
	}()
	return result

}

func main() {
	powerTow := raiseToPower(2)
	powerThree := raiseToPower(3)
	powerFour := raiseToPower(4)
	var in = 10.0
	fmt.Printf(fmt.Sprintf("%.2f ^ 2 = %.2f\n", in, powerTow(in)))
	fmt.Printf(fmt.Sprintf("%.2f ^ 3 = %.2f\n", in, powerThree(in)))
	fmt.Printf(fmt.Sprintf("%.2f ^ 4 = %.2f\n", in, powerFour(in)))

	numbers := make([]float64, 0, 10)
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Numbers are: ", numbers)
	fmt.Println("Numbers to ^2 are:", Apply(numbers, powerTow))
	fmt.Println("Numbers to ^3 are:", Apply(numbers, powerThree))
	fmt.Println("Numbers to ^4 are:", Apply(numbers, powerFour))
}
