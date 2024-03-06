package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	emp_id    int
	contactDetails
}

type contactDetails struct {
	emailId      string
	mobileNumber int
	address      string
}

func createNewStruct(firstName string, lastName string, emp_id int, emailId string, mobileNumber int, address string) employee {
	emp := employee{
		firstName: firstName,
		lastName:  lastName,
		emp_id:    emp_id,
		contactDetails: contactDetails{
			emailId:      emailId,
			mobileNumber: mobileNumber,
			address:      address,
		},
	}

	return emp
}

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

	jimmy := person{
		firstName: "Jimmy",
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

	/*
		The issue: Go is a pass by value language. Which means that when a value is passed to go, it makes a copy of that value, and that copy is made available to the code inside the function.
		Hence when we pass john to the updateName function, and try to update the firstname with the argument in the function, it will not update the value since it would have created a copy of the value in another memory.
		Solution: Use Pointers.
		1) We create a pointer to get the memory address of the struct, by adding an & operator preceeding to the struct person(&john). The & operator gives the address of the value that the variable is poining to. In this case it gives the address of john and stores the address in the johnPointer variable.
	*/
	johnPointer := &john
	fmt.Print("Before\n")
	john.print()
	fmt.Println("\n")
	jimmy.print()
	/*
		2) We then call the updateName function on the pointer rather than on the struct john, since calling it on john will create a copy of the value in the memory and wont modify the original struct
	*/
	johnPointer.updateName("Jimmy")
	fmt.Println("\nAfter\n")
	johnPointer.print()

	//john.updateName("Jimmy")
	//updateNewName(john)
	//john.print()

	//Lets try one example here
	sahil := createNewStruct("Sahil", "Kulkarni", 10023698, "sahil.kulkarni@quesscorp.com", 9538409954, "F4 Balaji Enclave 9th Cross")

	fmt.Println("\n")
	sahil.printName()
	fmt.Println("\n")
	sahil.print()

	sahil.updateEmployeeDetails("firstName", "Sameer")
	sahil.print()
}
