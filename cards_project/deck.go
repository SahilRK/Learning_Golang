package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

/*Steps we are going to follow to play with a deck of cards
1) Create a custom slice of cards called deck. Create a custom slice of strings called deck
2) Print the deck of cards using a receiver function
3) Create a deal function which returns two hands
4) Write two functions to Save the deck of cards to a file. This function consists of another function which will convert the deck of cards into a single string, which will be returned back.This function then writes back the string into a file using I/O Operations.
5) Get a deck of cards from the a file
*/

type deck []string

// When we want to return a value back, we add the return type after the function name
func newDeck() deck {
	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cardsDeck := []string{}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cardValue := value + " of " + suite
			cardsDeck = append(cardsDeck, cardValue)
		}
	}

	return cardsDeck

}

// Create a receiver function named print which prints all the cards. Receiver functions are instance functions that are executed on the custom datatype
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

//A function which accepts two arguments, deck and a handsize and splits the deck into two decks using the slicing

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// Convert the deck of strings into a single string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Write the single string to a file as a byte array
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Create a new deck from a file using the Readfile method
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// We have two options here when we want to deal with the error returned
		// #Option 1: log the error and return a new deck
		// #Option 2: log the error and entirely quit the program. We will be using this
		fmt.Println("Error: ", err)
		os.Exit(1)

	}

	strSlice := strings.Split(string(bs), ",")
	return deck(strSlice)
}

// Create a function to shuffle the deck of cards
func (d deck) shuffle() {
	//To generate a random number, we do the following.
	//1: Create a new source for the number from the current time
	//2: Generate a new rand type from that source
	//fmt.Println("Time now is: ", time.Now(), " and the Nano conversion of it is: ", time.Now().UnixNano())
	//newSource := rand.NewSource(time.Now().UnixNano())
	//newRandObj := rand.New(newSource)

	//We can ignore the second element which is the interator by just mentioning the index
	for ind := range d {
		//randNum := newRandObj.Intn(len(d) - 1)
		randNum := rand.Intn(len(d) - 1)
		fmt.Println("Random number is: ", randNum, " And value at Rand no is: ", d[randNum], " Index is: ", ind)

		d[ind], d[randNum] = d[randNum], d[ind]
	}
	fmt.Println(d)

}
