package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	// create new child goroutine for each link
	for _, link := range links {
		go checkLink(link, c)
	}

	// wait for chan to return some value, assign to l, and then iterate
	// pause for 5 seconds before creating new checkLink goroutine to not spam site
	for l := range c {
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	// if error, send status error msg through chan
	if err != nil {
		fmt.Println("Error:", link, "may currently be down.")
		c <- link
		return
	}

	// send status success msg through chan
	fmt.Println(link, "is up and visible!")
	c <- link
}
