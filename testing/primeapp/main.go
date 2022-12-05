package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()
	doneChannel := make(chan bool)
	go readUserInput(os.Stdin, doneChannel)
	// block until doneChannel gets a value
	<-doneChannel
	// Ok, now we can close the channel
	close(doneChannel)
	fmt.Println("Goodbye.")
}
func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		} else {
			fmt.Println(res)
			prompt()
		}
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number", false
	}
	_, msg := isPrime(numToCheck)
	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("-----------")
	fmt.Println("Enter a whole number, and we'll tell you " +
		"if it's a prime number or not. Enter q to quit")
	prompt()
}

func prompt() {
	fmt.Print("--> ")
}

func isPrime(number int) (bool, string) {
	if number == 0 || number == 1 {
		return false, fmt.Sprintf("%d is not prime number by definition!", number)
	}
	if number < 0 {
		return false, fmt.Sprintf("negative numbers are not prime by definition %d!", number)
	}
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", number, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number!", number)
}
