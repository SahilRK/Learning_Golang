package main

import "fmt"

// Create a custom type person of type struct. It is mandatory to mention the property names(keys) in the struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// Lets create another custom type called contact info, which will define the properties of contact
type contactInfo struct {
	phoneNumber int64
	email       string
	address     string
	pincode     int32
}

func main() {

	//Create and assign a new person which is of type struct. We then can embed a struct contactInfo within it
	john := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			phoneNumber: 9000100022,
			email:       "john.doe@doe.com",
			address:     "unknown",
			pincode:     560001,
		},
	}

	//Print a struct with the property names
	fmt.Printf("%+v", john)
	fmt.Printf("\n")

	//An alternate way of doing the above is
	var sahil person
	sahil.firstName = "Sahil"
	sahil.lastName = "Kulkarni"
	fmt.Printf("%+v", sahil)

}
