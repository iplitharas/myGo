package main

import (
	"fmt"
	"hello-world/hello"
	"hello-world/world"
	"io"
	"os"
)

func main() {
	fmt.Println("This is the main package")
	displayGreetings(os.Stdout)

}

func displayGreetings(w io.Writer) {
	_, err := fmt.Fprintf(w, hello.Greet())
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, world.Greet())
	if err != nil {
		return
	}

}
