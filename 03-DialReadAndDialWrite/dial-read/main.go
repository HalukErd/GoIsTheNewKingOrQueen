package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	//run 01-WriteToConn Then Run this
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	sl, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(sl))
}
