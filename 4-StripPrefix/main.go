package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/dog/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	err := http.ListenAndServe(":8080", nil)
	handleErr(err)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/resources/toby.jpg">
	`)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
