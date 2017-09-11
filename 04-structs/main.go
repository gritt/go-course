package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string

// same result
//	contact contactInfo
	contactInfo
}

func main() {

	jim := person{
		firstName: "jim",
		lastName:  "silva",

	    //contact : contactInfo {}

		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 96000,
		},
	}

	// print struct and attributes
	fmt.Printf("%+v", jim)
}
