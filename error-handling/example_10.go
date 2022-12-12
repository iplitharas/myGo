package main

import "fmt"

func doPanic(message string) {
	panic(message)
}

func doRecover(message string) {
	defer func() {
		fmt.Println("---Running defer---")
		if v := recover(); v != nil {
			fmt.Println("Value from recover!", v)
		}
	}()
	doPanic(message)
}

func main() {
	fmt.Println("Panic and recover demo")
	doRecover("Hi there!")
	fmt.Println("Recovered")
}
