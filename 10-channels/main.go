package main

import (
	"fmt"
	"net/http"
	"time"
)

// by default go tries to use one core, routines are handled by the go scheduler
// on one core then can run concurrently, on more cpus they are able to run truly in parallel

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
	// channels are ways of communication between routines / they're aware of routines
	// channels are typed, information sent through them must be typed
	for _, u := range l {
		go check(u, c)
	}

	// infinite loop
	// wait for the channel to return some value, when it return, assigns to i
	for i := range c {
		// receive from chanel, blocking call

		// not a great idea to pause the main routine,
		// as child routines will send data through the channel, a this main one cannot be sleeping
		//time.Sleep(5 * time.Second)
		//go check(i, c)

		// warning! when referencing a var which is being used / maintained by another go routine
		// we never ever attempt to reference the same var inside of two diff routines
		// to solve it, routines have to receive arguments by value, (if they'll use on startup time)

		// function literal / lambda / unnamed function - wrap some code to execute sometime in future
		go func(link string) {
			time.Sleep(5 * time.Second)
			check(link, c)
		}(i)

		// never try to access the same variable from a diff child routines
		// only share inf with a the child / new routine by passing it as arg, or communicating though the chanell
	}

	// same as
	// for {
	// 	// receive from chanel, blocking call
	// 	go check(<-c, c)
	// }
}

// signature block operation in the channel:
// channel given to READ/RECEIVE c <- chan
// channel given to WRITE/SEND c chan <-
func check(u string, c chan<- string) {

	_, err := http.Get(u)

	if err != nil {
		fmt.Println("this may be down", u)
		// send the url to the channel
		c <- u
		return
	}

	// send the url to the channel
	fmt.Println(u, "is up")
	c <- u
}
