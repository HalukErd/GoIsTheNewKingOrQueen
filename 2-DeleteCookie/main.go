package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/check", checkCuk)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	cukie := &http.Cookie{
		Name:  "H-Auth",
		Value: "Bearer Polar Bear",
		Path:  "/",
	}
	http.SetCookie(w, cukie)
	fmt.Println("New Cookie Created:", cukie)
	fmt.Fprintln(w, "Your Cookie:", cukie)
}

func checkCuk(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("H-Auth")
	fmt.Println("Cookie From Req:", cookie)
	fmt.Fprintln(w, "Your Cookie:", cookie)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("H-Auth")
	c.MaxAge = -1
	http.SetCookie(w, c)
}
