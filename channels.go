package main

import (
	"fmt"
	"net/http"
	"time"
)

func channels() {
	links := []string{
		"https://google.com",
		"https://planetoidi.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	/*
		for i := 0; i < 30; i++ {
			go checkLink(<-c, c)
		}
	*/

	for l := range c {
		// This is an anonymous function!
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	time.Sleep(5 * time.Second)
	if err != nil {
		fmt.Println(string("\033[31m"), "No response from ", link, "\033[0m")
		c <- link
		return
	}
	fmt.Println(string("\033[36m"), link, "русский военный корабль, иди на хуй", "\033[0m")
	c <- link
}
