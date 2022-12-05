package main

import (
	"bytes"
	"testing"
)

func Test_main(t *testing.T) {
	buf := new(bytes.Buffer)

	expectedOutput := "Helloworld!"
	displayGreetings(buf)
	result := buf.String()
	if result != expectedOutput {
		t.Fatalf("expected %s but got %s", expectedOutput, result)
	}
}
