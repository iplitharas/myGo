package main

import (
	"fmt"
	"strings"
)

func main() {
	inputText := ""
	fmt.Println("Enter a word:")
	fmt.Scan(&inputText)
	fmt.Println("You entered:", inputText)
	performTextCheck(inputText, "ian")
	testperformTextCheck()

}

// Given an input string: `textForSearch`  and a string containing the `chars` for searching `searchPatterns`
// returns true if:
// a)the first char of the  `textForSearch` equals with the first of `searchPatterns`
// b)the last char of the  `textForSearch` equals with the last of `searchPatterns`
// c)the `textForSearch` contains the second pattern from `searchPatterns`
func performTextCheck(textForSearch string, searchPatterns string) int {
	textForSearch = strings.ToLower(textForSearch)
	check := strings.Contains(textForSearch, string(searchPatterns[1])) &&
		textForSearch[0] == searchPatterns[0] &&
		textForSearch[len(textForSearch)-1] == searchPatterns[2]
	switch check {
	case true:
		{
			fmt.Println("Found!")
			return 1
		}
	case false:
		{
			fmt.Println("Not Found")
		}

	}
	return -1

}

func testperformTextCheck() {
	searchPattern := "ian"
	//Positive test cases
	performTextCheck("ian", searchPattern)
	performTextCheck("iaaaan", searchPattern)
	performTextCheck("iuiygaygn", searchPattern)
	performTextCheck("Ian", searchPattern)
	performTextCheck("I d skd a efju N", searchPattern)

	// Negative test cases
	performTextCheck("“ihhhhhn", searchPattern)
	performTextCheck("“ina", searchPattern)
	performTextCheck("“xian", searchPattern)

}
