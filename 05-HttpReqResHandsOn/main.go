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

	m, uri := request(conn)

	response(conn, m, uri)
}

func request(conn net.Conn) (string, string) {
	i := 0
	var methodType string
	var uri string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fields := strings.Fields(ln)
		if i == 0 {
			methodType = fields[0]
			uri = fields[1]
			fmt.Println("Method Type:" + methodType)
		}
		if i == 11 && methodType == "GET" {
			//if ln != "" && strings.Fields(ln)[0] == "Referer:" && methodType == "GET" {
			url := fields[1]
			fmt.Println("Url:", url)
		}
		if ln == "" {
			break
		}
		i++
	}
	return methodType, uri
}

func response(conn net.Conn, methodType string, uri string) {
	if methodType == "GET" && uri == "/" {
		homeGetRes(conn)
	}
	if methodType == "GET" && uri == "/login" {
		loginGetRes(conn)
	}
}

func loginGetRes(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Please Login</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func homeGetRes(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
