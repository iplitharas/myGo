package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Create and write to a file example..")

	file, err := os.Create("newFile.txt")
	if err != nil {
		fmt.Printf("error while trying to create the file %q\n", err)
	}
	// Create a buffer
	outputWriter := bufio.NewWriter(file)
	fileContents := []string{"Hello there\n", "Can I create files?"}
	for _, content := range fileContents {
		_, err := outputWriter.WriteString(content)
		if err != nil {
			return
		}
	}
	// write everything to the file
	outputWriter.Flush()
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

}
