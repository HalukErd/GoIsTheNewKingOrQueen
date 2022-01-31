package main

import (
	"html/template"
	"log"
	"net/http"
)

type customer struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

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
	//v := req.FormValue("q")

	fName := req.FormValue("first")
	lName := req.FormValue("last")
	sub := req.FormValue("subscribe") == "on"

	data := customer{
		FirstName:  fName,
		LastName:   lName,
		Subscribed: sub,
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
