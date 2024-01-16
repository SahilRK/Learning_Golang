package main

import "fmt"

// Create a custom type person of type struct. It is mandatory to mention the property names(keys) in the struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

// Lets create another custom type called contact info, which will define the properties of contact
type contactInfo struct {
	phoneNumber int64
	email       string
	address     string
	pincode     int32
}

// Create a receiver function print
func (p person) print() {
	fmt.Printf("%+v", p)
}

//Create a receiver function to update a persons details

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
