package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panicf("error trying to establish the connection: ", err)
	}

	clientRequest(conn)
	fmt.Fprintf(conn, "Hey, there")

}

func clientRequest(conn net.Conn) {
	// User input
	keyboardScanner := bufio.NewReader(os.Stdin)
	// connection scanner
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		fmt.Printf("[Server]: %q\n", text)
		userInput, _ := keyboardScanner.ReadString('\n')
		fmt.Fprintf(conn, userInput)
	}

	defer conn.Close()

}
