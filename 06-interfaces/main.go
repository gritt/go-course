package main

import "fmt"

// go automatically matches interfaces and type methods
// do not have to explicitly declare that the type implements an interface
type bot interface {
	getGreeting() string
}
type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sp := spanishBot{}


	// both bots use bot interface
	printGreeting(eb)
	printGreeting(sp)
}

// receive the interface as type
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// omitting the value "eb" as var will not be used
func (englishBot)getGreeting() string {
	// let's suppose this code is very specific / custom implementation
	// cant be reused
	return "Hi There"
}

// omitting the value "eb" as var will not be used
func (spanishBot)getGreeting() string {
	// let's suppose this code is very specific / custom implementation
	// cant be reused
	return "Hola"
}