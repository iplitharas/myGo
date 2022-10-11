package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)

}

func main() {

	var wg = sync.WaitGroup{}
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"pi",
		"zeta",
		"epsilon",
		"theta",
	}
	wg.Add(len(words))
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	// because the go routine runs so fast
	//if we want to see the first goroutine we have to wait
	//time.Sleep(1 * time.Second)
	// instead we use wg

	wg.Wait()
}
