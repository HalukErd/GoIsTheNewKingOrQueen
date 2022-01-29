package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	i := 0
	var uri string
	var method string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if len(ln) != 0 && i == 0 {
			fields := strings.Fields(ln)
			uri = fields[1]
			method = fields[0]
		}
		if ln == "" {
			break
		}
		i++
	}
	response(conn, method, uri)

}

func response(conn net.Conn, method string, uri string) {

	if method == "GET" && uri == "/" {
		home(conn)
	}
	if method == "GET" && uri == "/apply" {
		home(conn) // different method
	}
	if method == "POST" && uri == "/apply" {
		home(conn) // another different method
	}
}

func home(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	fmt.Fprintf(conn, body)
}
