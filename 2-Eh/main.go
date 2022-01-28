package main

import (
	"fmt"
	"io"
	"net/http"
)

type login int
type signup int

func (h login) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hey hey")
	io.WriteString(w, "Hey hey2\n")
	fmt.Fprintln(w, "Please Login")
}

func (s signup) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hey hey")
	io.WriteString(w, "Hey hey2\n")
	fmt.Fprintln(w, "Please Register")
}

func main() {
	var s signup
	var l login
	mux := http.NewServeMux()

	mux.Handle("/login", l)
	mux.Handle("/register", s)
	
	http.ListenAndServe(":8080", mux)
}
