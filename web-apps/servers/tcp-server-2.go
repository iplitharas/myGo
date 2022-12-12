package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(con)
	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println(err)
	}
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()
	fmt.Println("Connection expired")
}
