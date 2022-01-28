package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var m = make(map[string]string)

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
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")
	//conn.SetDeadline(time.Now().Add(10 * time.Second))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		ln := scanner.Text()
		sl := strings.Fields(ln)
		res := exeCmd(sl)
		fmt.Printf("Req: %s Fields:%v Response:%s", ln, sl, res)
		fmt.Fprintln(conn, res)
	}

	fmt.Println("Code got here")
}

func exeCmd(sl []string) string {

	switch sl[0] {
	case "GET":
		return m[sl[1]]
	case "SET":
		m[sl[1]] = sl[2]
		return sl[1] + " is created."
	case "DEL":
		delete(m, sl[1])
		return sl[1] + " is deleted."
	default:
		return "Invalid Command"
	}
}
