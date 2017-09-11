package main

import "fmt"

func main() {

	colors := map[string]string {
		"red": "#ff0000",
		"white": "#ffffff",
		"gray": "#cccccc",
	}

	printMap(colors)
}

// has to know the map structure
func printMap(c map[string]string)  {
	for color, hex := range c {
		fmt.Println("hexcode for color", color, "is", "hex", hex)
	}
}
