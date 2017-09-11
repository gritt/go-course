package main

import "fmt"

func main() {

	// define a empty map, where all keys will be int, and values will be string
	//var colors map[int]string

	// A new, empty map value is made using the built-in function make, which takes the map type and an optional capacity hint as arguments:
	//
	// make(map[string]int)
	// make(map[string]int, 100)
	// The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, with the exception of nil maps. A nil map is equivalent to an empty map except that no elements may be added.

	// another syntax
	colors := make(map[int]string)

	// another syntax
	//colors := map[int]string {
	//	9: "#ff0000",
	//	2: "#ffffff",
	//}

	// add to map
	colors[10] = "#yellow"

	// delete from map
	delete(colors, 10)

	fmt.Println(colors)
}
