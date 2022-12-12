package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	// close the port
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleTCP(conn)

	}

}

func handleTCP(conn net.Conn) {
	fmt.Fprintf(conn, "Hello, I'm listening... Ask me anything you like\n")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("[Client]: %q\n", scanner.Text())
		fmt.Fprintf(conn, "%s\n", strings.ToUpper(text))
	}
	//// Close the connection
	defer conn.Close()

}
