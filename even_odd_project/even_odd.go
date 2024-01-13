package main

import "fmt"

type intList []int

func newIntList() intList {
	listOfNum := intList{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	return listOfNum
}

func (iL intList) evenOrOddd() {
	for _, val := range iL {
		if val%2 == 0 {
			fmt.Println(val, " is an even number")
		} else {
			fmt.Println(val, " is an odd number")
		}
	}
}
