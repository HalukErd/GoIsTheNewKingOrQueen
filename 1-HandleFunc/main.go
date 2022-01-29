package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Better than cat")
}

func name(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.Form.Get("name")
	fmt.Fprintln(w, "Are you", name, "?")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", name)
	http.ListenAndServe(":8080", nil)
}
