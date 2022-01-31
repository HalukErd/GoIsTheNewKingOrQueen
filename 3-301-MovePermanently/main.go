package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.FormValue("fname"))
	fmt.Println("Request Method:", req.Method, "URL:", req.URL)

	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Method:", req.Method, req.URL)
}
