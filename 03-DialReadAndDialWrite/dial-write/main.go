package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//run 01-WriteToConn Then Run this
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(conn, "Dialed you!")
}
