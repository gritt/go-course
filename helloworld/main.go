package main

import "fmt"

func main() {

	//explicitly define type

	var card string = "Ace of Spaces"

	// go infer the type, cannot reassign with :=
	//card := "Ace of spaces"

	fmt.Println(card)

	fmt.Println("Hello World")
}
