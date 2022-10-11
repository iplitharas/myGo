package main

import "fmt"

func main() {
	var firstInt int
	var secondInt int
	parseInput(&firstInt, &secondInt)
	fmt.Println("You entered two numbers:\n"+
		"First number is:", firstInt, ""+
		"and second is:", secondInt)

}

func parseInput(x *int, y *int) {
	fmt.Println("Enter two float numbers (space separated):")
	_, _ = fmt.Scan(x, y)

}
