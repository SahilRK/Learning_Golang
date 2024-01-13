package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected to return 53 cards, but it returned %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected to return Ace of Spades, but returned %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected to return King of Clubs, but returned %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected to return 52 cards, but returned %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
