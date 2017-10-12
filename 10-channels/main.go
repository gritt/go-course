package main

import (
	"fmt"
	"net/http"
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

	// channel to share strings between routines
	c := make(chan string)

	// because of the routine_bug.png, we use channels
	// channels are ways of communication bewteeen routines / they're aware of routines
	// channels are typed, information sent through them must be typed
	for _, u := range l {
		go check(u, c)
	}

	// infinite loop
	for {
		// receive from chanell, blocking call
		go check(<-c, c)
	}
}

// this is the "blocking call",
// it freezes the for loop (caller) till it get a response
func check(u string, c chan string) {
	_, err := http.Get(u)
	if err != nil {
		fmt.Println("this may be down", u)

		// send the url to the channel
		c <- u
		return
	}

	fmt.Println(u, "is up")
	c <- u
}
