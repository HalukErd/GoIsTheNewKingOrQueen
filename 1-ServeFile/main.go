package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	err := http.ListenAndServe(":8080", nil)
	handleErr(err)

}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/toby.jpg">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	handleErr(err)
	defer f.Close()
	io.Copy(w, f)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
