package main

import "fmt"

// go automatically matches interfaces and type methods
// do not have to explicitly declare that the type implements an interface
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

// [1] make any function that returns a string compatible with botGenerator
type botGenerator func() string

// [2] AND, receives the `outdated` function and run in wrapped in getGreeting, making it now an implementation of bot interface
func (p botGenerator) getGreeting() string {
	return p()
}

// THIS -> is the outdated function that needs to be "wrapped" to be compatible with the needed interface (bot interface)
func getBomDia() string {
	return "Bom dia"
}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(botGenerator(getBomDia))

	// both bots use bot interface
	printGreeting(eb)
	printGreeting(sp)

}

// receive the interface as type
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// omitting the value "eb" as var will not be used
func (englishBot) getGreeting() string {
	// let's suppose this code is very specific / custom implementation
	// cant be reused
	return "Hi There"
}

// omitting the value "eb" as var will not be used
func (spanishBot) getGreeting() string {
	// let's suppose this code is very specific / custom implementation
	// cant be reused
	return "Hola"
}
