package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("assets/*.gohtml"))
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.Handle("/resources", http.StripPrefix("resources", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dog(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
