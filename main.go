package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}
	c := make(chan string)

	//for _, link := range links {
	//	checkLink(link, c)
	//}

	for _, link := range links {
		go checkLink(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//	log.Println(<-c)
	//}
	for l := range c {
		//time.Sleep(10 * time.Second)
		//go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

	//time.Sleep(5 * time.Second)
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		log.Println(l, "is down.")
		c <- l
		return
	}
	log.Println(l, "is up.")
	c <- l
}
