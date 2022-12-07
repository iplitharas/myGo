package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter something")
	inputReader := bufio.NewReader(os.Stdin)
	userInput, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("You typed: %q\n", userInput)
}
