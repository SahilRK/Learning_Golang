package main

import "fmt"

// Create a custom type person of type struct. It is mandatory to mention the property names(keys) in the struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

// EMBEDDING STRUCT
// Lets create another custom type called contact info, which will define the properties of contact. We will be embedding this inside the person struct.
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
/*
2) We then call the updateName function on the pointer rather than on the struct john, since calling it on john will create a copy of the value in the memory and wont modify the original struct.

The * operator tells Go to fetch the value that the pointer variable is pointing to in the memory.
*/

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

/*
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
*/
func (emp employee) printName() {
	fmt.Println("Emp First name is: ", emp.firstName, " and Last name is: ", emp.lastName)
}

func (emp employee) print() {
	fmt.Printf("%+v")
}

func (empPointer *employee) updateEmployeeDetails(propertyName string, propertyValue string) {
	propertyName1 := propertyName
	(*empPointer).propertyName1 = propertyValue
}
