package main

import (
	"fmt"
	"os"
)

type slice []int

func (s *slice) add(element int) {
	*s = append(*s, element)
}

func main() {

	var numbers slice
	fmt.Println(numbers)
	numbers.add(5)
	numbers.add(5)
	fmt.Println(numbers)
}

func readNumbers(numbers []int, total int) {
	var userInput int
	for i := 0; i < total; i++ {
		fmt.Printf("Enter the #%d number: ", i+1)
		_, err := fmt.Scan(&userInput)
		numbers = append(numbers, userInput)
		if err != nil {
			fmt.Println("Error during input:", err)
			os.Exit(1)
		}

	}

}
