package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

var tpl *template.Template

type hndlr string

func (h hndlr) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("H-Authorization", "Bearer Polar Bear")
	w.Header().Set("content-type", "text/html")
	data := struct {
		Method        string
		Url           *url.URL
		Submissions   url.Values
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		Method:        req.Method,
		Url:           req.URL,
		Submissions:   req.Form,
		Header:        req.Header,
		Host:          req.Host,
		ContentLength: req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)

}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h hndlr
	http.ListenAndServe(":8080", h)
}
