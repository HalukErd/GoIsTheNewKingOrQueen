package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		io.WriteString(conn, "\nHello Tcp Server\n")
		fmt.Fprintln(conn, "You're my daily now")
		fmt.Fprintln(conn, "So today was yesterday")

		conn.Close()
	}

}
