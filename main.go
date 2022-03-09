package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Ohh, Elden Ring")
	if err != nil {
		return
	}
}
