package main

/*
type book string

func (b book) printline(desc string) string {
	return string(b) + " " + desc
}

func (d deck) printLine() string {

}
*/

import (
	"fmt"
)

func printNum() {
	var strSlice []string
	for i := 1; i <= 100; i++ {
		strSlice = append(strSlice, fmt.Sprintf("%d", i))
	}
	fmt.Println(strSlice)
}
