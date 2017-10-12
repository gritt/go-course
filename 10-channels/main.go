package main

import (
	"fmt"
	"net/http"
	"os"
)

// by default go tries to use one core, routines are handled by the go scheduler
// on one core then can run concurently, on more cpus they are able to run trully in parallel

// scheduler run one thread on each logical core
// this behaviour can be changed
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
func check(u string) {
	_, err := http.Get(u)
	if err != nil {
		fmt.Println("this may be down", u)
		return
	}

	fmt.Println(u, "is up")
}
