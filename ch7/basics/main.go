package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " might be down.")
	} else {
		fmt.Println(link + " is up.")
	}

	c <- link
}

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://golang.org",
		"https://facebook.com",
		"https://stackoverflow.com",
	}
	c := make(chan string, 2)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}
