package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Reading from file example..")
	readFileUsingOpen()
	//readfileUsingOpenFile()

}

func readFileUsingOpen() {
	fmt.Println("reading file using 'os.Open' ")
	inputFile, inputError := os.Open("input.txt")
	if inputError != nil {
		fmt.Println(inputError)
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {

		}
	}(inputFile)

	readUsingBufNewReader(inputFile)
}

func readfileUsingOpenFile() {
	fmt.Println("reading file using `os.OpenFile`")
	inputFile, inputError := os.OpenFile("input.txt", os.O_RDONLY, 0666)
	if inputError != nil {
		fmt.Println(inputError)
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {

		}
	}(inputFile)
	readUsingBufferRead(inputFile)

}

func readUsingBufNewReader(inputFile *os.File) {
	inputReader := bufio.NewReader(inputFile)
	for {
		fileLine, err := inputReader.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Print(fileLine)
	}
}

func readUsingBufferRead(inputFile *os.File) {
	buf := make([]byte, 1024)
	inputReader := bufio.NewReader(inputFile)

	for {
		n, err := inputReader.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(string(buf))
}
