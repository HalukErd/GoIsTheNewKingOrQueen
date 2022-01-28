package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

type hndlr string

func (h hndlr) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)

}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h hndlr
	http.ListenAndServe(":8080", h)
}
