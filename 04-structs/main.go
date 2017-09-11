package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {

	// populate strings
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	// attributes are defined, but nil
	var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"


	// print struct and attributes
	fmt.Printf("%+v", alex)
}
