package main

import (
	"os"
	"testing"
)

//ensure that you have correct number of cards in your deck
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
	//ensure that the first card is Ace Of Spades
	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected first card Ace Of Spades, but got %v", d[0])
	}
	//ensure that the first card is Four Of Clubs
	if d[len(d)-1] != "Four Of Clubs" {
		t.Errorf("Expected last card Four Of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFIle("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
