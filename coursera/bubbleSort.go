package main

/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation.
Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.



*/
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the size the array: (up to 10)")
	var totalNumbers int
	_, err := fmt.Scan(&totalNumbers)
	if err != nil {
		fmt.Println("Error during input:", err)
		os.Exit(1)
	}
	if totalNumbers > 10 {
		fmt.Println("Maximum size should be 10\nExiting..")
		os.Exit(1)

	}
	numbers := parseInput(totalNumbers)
	fmt.Println("Before sorting:", numbers)
	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)

}

func parseInput(totalNumbers int) []int {
	var userInput int
	numbers := make([]int, 0, 10)
	fmt.Printf("Enter in total #%d numbers\n", totalNumbers)
	for i := 0; i < totalNumbers; i++ {
		fmt.Printf("Enter the #%d number: ", i+1)
		_, err := fmt.Scan(&userInput)
		numbers = append(numbers, userInput)
		if err != nil {
			fmt.Println("Error during input:", err)
			os.Exit(1)
		}
	}
	return numbers

}

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func Swap(numbers []int, index int) {
	numbers[index], numbers[index+1] = numbers[index+1], numbers[index]
}
