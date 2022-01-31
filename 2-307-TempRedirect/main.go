package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func bar(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Location", "/")
	//w.WriteHeader(http.StatusSeeOther)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)

	fmt.Println(req.FormValue("fname"))
	fmt.Println("Request Method:", req.Method, "URL:", req.URL)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Body:", req.FormValue("fname"))
	fmt.Println("Request Method:", req.Method, "URL:", req.URL)

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleErr(w, err)
}

func foo(w http.ResponseWriter, req *http.Request) {
	//v := req.FormValue("q")
	fmt.Println("Request Method:", req.Method, req.URL)
}

func handleErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
