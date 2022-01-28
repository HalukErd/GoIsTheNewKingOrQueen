package main

import (
	"fmt"
	"io"
	"net/http"
)

type hndlr int

func (h hndlr) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hey hey")
	io.WriteString(w, "Hey hey2\n")
	switch req.URL.Path {
	case "/login":
		fmt.Fprintln(w, "Please Login")
	case "/register":
		fmt.Fprintln(w, "Please Sign Up")
	}
}

func main() {
	var h hndlr
	http.ListenAndServe(":8080", h)
}
