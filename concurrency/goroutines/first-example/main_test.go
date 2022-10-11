package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

// test  `printSomething` as a goroutine
func Test_printSomething(t *testing.T) {
	// Test Setup
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	//Specify the `Stdout` to write at our Pipe`
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go printSomething("epsilon", &wg)
	wg.Wait()

	//close Pipe
	w.Close()
	// and read from it
	result, _ := io.ReadAll(r)
	output := string(result)

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon, but it's not there")
	}
	os.Stdout = stdOut
}
