package main

import (
	"fmt"
	"io"
	"net/http"
)

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hey hey")
	io.WriteString(w, "Hey hey2\n")
	fmt.Fprintln(w, "Please Login")
}

func signup(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hey hey")
	io.WriteString(w, "Hey hey2\n")
	fmt.Fprintln(w, "Please Register")
}

func main() {

	http.HandleFunc("/login", login)
	http.HandleFunc("/register", signup)

	http.ListenAndServe(":8080", nil)
}
