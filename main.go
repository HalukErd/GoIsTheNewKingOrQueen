package main

import (
	"fmt"
	"net/http"
)

type hndlr string

func (h hndlr) ServeHTTP(resWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resWriter, "Here is the place you should code")
}

func main() {
	var h hndlr
	http.ListenAndServe(":8080", h)
}
