// The main package is declared to create an executable type package
package main

func main() {
	/* cards := newDeck()

	cards.print()

	firstDeal, secondDeal := deal(cards, 5)

	fmt.Println(firstDeal)
	fmt.Println("------------------------------------------")
	fmt.Println(secondDeal)

	cards.saveToFile("mycards.txt") */

	//cards := newDeckFromFile("mycards.txt")
	//fmt.Println(cards)
	//cards.shuffle()

	printNum()

}

/* func strToByte() {
	testStr := "Hi there"
	fmt.Println([]byte(testStr))
}
*/
/*//Import external packages

// The main function is created since we want to create an executable package, which is required if we have a main package
func main() {

	cards := newDeck()

	newBook := book("Rich dad Poor dad")
	fmt.Println(newBook)

	fmt.Println(newBook.printline(("is a great book")))

	//fmt.Println(cards)

	//When we loop through the Datastructure using a for loop, we use the := operator, since for each loop, the variables are removed and redeclared and reinitialised
	//We also use the range keyword to loop through the values in the datastructure.
	for indx, card_val := range cards {
		fmt.Println(indx, card_val)
	}

	//We call the print function from the cards variable
	//cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("-----------------------------------")
	remainingCards.print()
	//------------ Practice section ----------------------//
	//var newBook book = "Easy ways to learn Golang"
	//newBook.printline()
}

// When we create a function and want to return a value, we have to mention the return datatype after the function
func newCard() string {
	return "Queen of Hearts"
}

*/
