package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	// Given a new string
	testingStr := "Testing str"
	var wg sync.WaitGroup
	wg.Add(1)
	//When
	//I call the `updateMessage` as a goroutine
	go updateMessage(testingStr, &wg)
	wg.Wait()
	//Then
	// I'm expecting the msg to be updated
	if !strings.Contains(msg, "Testing str") {
		t.Errorf("Error ")
	}
}

func Test_printMessage(t *testing.T) {
	// store current os.Stdout
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	// Given a new string
	msg = "Testing Str"
	// When I call the `printMessage`
	printMessage()
	// Then I'm expecting the msg to be printed
	// in the stdOut
	err := w.Close()
	if err != nil {
		return
	}
	result, _ := io.ReadAll(r)
	output := string(result)
	if !strings.Contains(output, "Testing Str") {
		t.Errorf("Expecting: %s Received: %s", msg, result)
	}
	// Reset os.Stdout
	os.Stdout = stdOut

}
