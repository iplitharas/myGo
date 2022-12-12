package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// Sanity test: `telnet localhost 8080`

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicf("cannot open a port due to: %s", err)
	}
	defer func(li net.Listener) {
		err := li.Close()
		if err != nil {

		}
	}(li)
	for {
		con, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(con, "\nHello from TCP server\n")
		fmt.Fprintln(con, "How is your day?")
		fmt.Fprintf(con, "%v", "How is your day?")
		con.Close()
	}
}
