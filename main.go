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

	for _, link := range links {
		go checkLink(link, c)
	}

	/*
	  When we receive a message through chan, that new value is assigned to l
	  We pass l off to the function literal and that string is copied in memory
	  The goroutine has access to that copy instead of the original value of l
	*/
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Error:", link, "may currently be down.")
		c <- link
		return
	}

	fmt.Println(link, "is up and visible!")
	c <- link
}
