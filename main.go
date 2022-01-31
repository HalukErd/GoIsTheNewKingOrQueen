package main

import (
	"fmt"
	"net/http"
)

var m = make(map[string]int)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/check", checkCuk)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	cukie := &http.Cookie{
		Name:  "H-Auth",
		Value: req.RemoteAddr,
		Path:  "/",
	}
	http.SetCookie(w, cukie)
	fmt.Println("New Cookie Created:", cukie)
	fmt.Fprintln(w, "Your Count:", cukie)
}

func checkCuk(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("H-Auth")
	if err == http.ErrNoCookie {
		cukie := &http.Cookie{
			Name:  "H-Auth",
			Value: req.RemoteAddr,
			Path:  "/",
		}
		http.SetCookie(w, cukie)
	}
	m[cookie.Value] = m[cookie.Value] + 1
	fmt.Println("Cookie From Req:", cookie)
	fmt.Fprintln(w, "Your Count:", m[cookie.Value])
}
