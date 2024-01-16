package main

func main() {

	//Create and assign a new person which is of type struct. We then can embed a struct contactInfo within it
	john := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			phoneNumber: 9000100022,
			email:       "john.doe@doe.com",
			address:     "unknown",
			pincode:     560001,
		},
	}

	/* //Print a struct with the property names
	fmt.Printf("%+v", john)
	fmt.Printf("\n")
	fmt.Println(john.contactInfo.email)

	//An alternate way of doing the above is
	var sahil person
	sahil.firstName = "Sahil"
	sahil.lastName = "Kulkarni"
	fmt.Printf("%+v", sahil) */

	john.print()
	john.updateName("Jimmy")

}
