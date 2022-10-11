package main

import (
	"fmt"
	"sync"
)

var msg string

func main() {
	fmt.Println("Race conditions example 1")
	msg = "hello world"
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(2)
	//go UpdateWithRaceCondition("Hello universe", &wg)
	//go UpdateWithRaceCondition("Hello Race conditions", &wg)
	go UpdateWithMutex("Hello universe", &wg, &mutex)
	go UpdateWithMutex("Bye bye Race conditions", &wg, &mutex)
	wg.Wait()
	fmt.Println("Input is: ", msg)

}
func UpdateWithRaceCondition(newMessage string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Entering UpdateWithRaceCondition with: %s\n", msg)
	msg = newMessage

}

func UpdateWithMutex(newMessage string, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	fmt.Printf("Entering UpdateWithMutex with: %s\n", msg)
	msg = newMessage
	m.Unlock()
}
