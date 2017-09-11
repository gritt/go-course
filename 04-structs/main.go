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
		firstName: "Jim",
		lastName:  "Haskel",

		//contact : contactInfo {}

		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 96000,
		},
	}

	jim.updateName("Joe")

	// another way to deal
	// creates a pointer to the mem location of jim
	jimPointer := &jim
	jimPointer.updateLastName("Doe")


	jim.print()
}

// pointer to change the value of the actual person instance (mem address value)
// without pointer, go creates a copy of the struct p to another mem block
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateLastName(newLastName string)  {
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// turn address into value with *address
// turn value into address with &value
