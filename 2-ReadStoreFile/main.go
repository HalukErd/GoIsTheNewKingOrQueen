package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		handleErr(w, err)
		bs, err := io.ReadAll(f)
		handleErr(w, err)

		fmt.Println(h.Filename, h.Header, ":::", string(bs))

		dst, err := os.Create(path.Join("./user/", h.Filename))
		_, err = dst.Write(bs)
		handleErr(w, err)
		fmt.Println(bs)

	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleErr(w, err)
}

func handleErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
