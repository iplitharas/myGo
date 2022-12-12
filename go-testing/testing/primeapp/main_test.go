package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	testingInput := 0
	result, msg := isPrime(testingInput)
	if result {
		t.Errorf("with %d as test parameter got true, but expected false", testingInput)
	}
	if msg != "0 is not prime number by definition!" {
		t.Errorf("wrong message returned: %s", msg)
	}

}

func Test_isPrimeWithTableTest(t *testing.T) {
	// Testing params
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"prime", 3, true, "3 is a prime number!"},
		{"not prime", 4, false, "4 is not a prime number because it is divisible by 2!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"not prime", 12, false, "12 is not a prime number because it is divisible by 2!"},
		{"negative", -4, false, "negative numbers are not prime by definition -4!"},
		{"negative", 0, false, "0 is not prime number by definition!"},
		{"negative", 1, false, "1 is not prime number by definition!"},
	}
	for _, e := range primeTests {
		// Given a number
		// When I call the `isPrime`
		result, msg := isPrime(e.testNum)
		// Then I should expect the right
		// result and message
		if e.expected && !result {
			t.Errorf("Test:%s\nexpected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("Test: %s\nexpected false but got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("Test:%s\n expected: %s\n but got: %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of the `os.Stdout`
	oldOut := os.Stdout
	//create a read and write pipe
	r, w, _ := os.Pipe()
	// set `os.Stdout` to our write pipe
	os.Stdout = w
	prompt()
	// close our writer
	_ = w.Close()
	// reset our `Stdout`
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)
	// perform our test
	if string(out) != "--> " {
		t.Errorf("expected -->\nreceived: %s", out)
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	_ = w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)
	// perform our test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("Intro text not correct got: %s", string(out))

	}
}

func Test_checkNumbers(t *testing.T) {
	inputTests := []struct {
		testName        string
		userInput       string
		closeChannel    bool
		expectedMessage string
	}{
		{"prime", "7", false, "7 is a prime number!"},
		{"negative number", "0", false, "0 is not prime number by definition!"},
		{"invalid use input", "number", false, "Please enter a whole number"},
		{"exit with lower q ", "q", true, ""},
		{"exit with capital Q ", "Q", true, ""},
	}
	for _, testInput := range inputTests {
		input := strings.NewReader(testInput.userInput)
		reader := bufio.NewScanner(input)
		res, closeChannel := checkNumbers(reader)
		if !strings.EqualFold(res, testInput.expectedMessage) {
			t.Errorf("Test case:%s\n,Expected: %s\nreceived %s",
				testInput.testName, testInput.expectedMessage, res)
		}
		if closeChannel != testInput.closeChannel {
			t.Errorf("Test case:%s\nExpected %v, received %v",
				testInput.testName,
				testInput.closeChannel,
				closeChannel)
		}

	}

}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)
	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))
	// io.Reader is an interface
	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
