package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Better than cat")
}

func name(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.Form.Get("name")
	tpl.ExecuteTemplate(w, "index.gohtml", name)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", name)
	http.ListenAndServe(":8080", nil)
}
