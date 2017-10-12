package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// website status checker

	l := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	for _, u := range l {
		check(u)
	}

	os.Exit(1)
}

// this is the "blocking call",
// it freezes the for loop (caller) till it get a response
// to fix it, this whole logic can be executed in a new routine, like a thread
func check(u string) {
	_, err := http.Get(u)
	if err != nil {
		fmt.Println("this may be down", u)
		return
	}

	fmt.Println(u, "is up")
}
