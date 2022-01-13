package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Println(resp)

	//var bs byte[]
	//bs := []byte{}

	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	log.Println(string(bs))
	log.Printf("This response has %v bytes", len(bs))
	return len(bs), nil
}
