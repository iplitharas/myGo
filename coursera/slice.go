package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	userInput := ""
	numbers := make([]int, 0)
	for strings.ToLower(userInput) != "X" {
		fmt.Println("Enter an integer or type `X` for exit..")
		_, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Println("An error occurred:", err)
			return
		}
		integer, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("This is not an integer...\nerror is:", err)
			return
		} else {
			numbers = append(numbers, integer)
			sort.Sort(sort.IntSlice(numbers))
			fmt.Println("Numbers after sorting are: ", numbers)
		}
	}
}
