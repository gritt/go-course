package main

import "fmt"

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sp := spanishBot{}


	//todo, use interfaces to make reuse of printGreeting
	printGreeting(eb)
	printGreeting(sp)
}

// cannot overload functions, even they have different signature (args)
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sp englishBot)  {
	fmt.Println(sp.getGreeting())
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