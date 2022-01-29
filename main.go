package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprintln(w, "Hey hey")
	io.WriteString(w, "Hey hey2\n")
	fmt.Fprintln(w, "Please Login")
}

func signup(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprintln(w, "Hey hey")
	io.WriteString(w, "Hey hey2\n")
	fmt.Fprintln(w, "Please Register")
}

func loginProcess(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	uid := params.ByName("username")
	fmt.Fprintln(w, "Welcome", uid)
}

func main() {
	rot := httprouter.New()
	rot.GET("/login", login)
	rot.GET("/register", signup)
	rot.GET("/login/:username", loginProcess)

	log.Fatal(http.ListenAndServe(":8080", rot))
}
