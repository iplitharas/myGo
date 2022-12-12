package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Listening at: http://localhost:8080")
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
		}
		go parseRequest(conn)

	}

}

func parseRequest(conn net.Conn) {
	defer conn.Close()
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		// first line is the method
		if i == 0 {
			fmt.Println(ln)
			mux(conn, ln)

		}
		if ln == "" {
			break
		}
		i++
	}

}

func mux(conn net.Conn, line string) {
	m := strings.Fields(line)[0]
	uri := strings.Fields(line)[1]
	fmt.Printf("Requested method is: %q\n", m)
	fmt.Printf("Requested uri is: %q\n", uri)

	if m == "GET" && uri == "/index" {
		indexHandler(conn)
	} else {
		inProgress(conn, m, uri)
	}
}

func indexHandler(conn net.Conn) {
	fmt.Println("Entering indexHandler")
	htmlResponse := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Global functions</title>
</head>
<body>
<h1>
Index home page
</h1>
</body>
</html>`
	// HTTP response
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(htmlResponse))
	fmt.Fprint(conn, "Content-Type text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintf(conn, htmlResponse)

}

func inProgress(conn net.Conn, method, uri string) {
	fmt.Println("Entering indexHandler")
	htmlResponse := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Global functions</title>
</head>
<body>
<h1>
<ul>
<li> Requested Method: {{ .Method }} </li>
<li> URI {{ .Uri }} in progress </li>
</ul>
</h1>
</body>
</html>`
	// HTTP response
	tmpl := template.New("my_template")
	tmpl.Parse(htmlResponse)

	type requestInfo struct {
		Method string
		Uri    string
	}
	req := requestInfo{method, uri}
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(htmlResponse))
	fmt.Fprint(conn, "Content-Type text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	//fmt.Fprintf(conn, htmlResponse)
	err := tmpl.ExecuteTemplate(conn, "my_template", req)
	if err != nil {
		log.Println(err)
	}

}
