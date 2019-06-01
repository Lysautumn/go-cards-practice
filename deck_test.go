package main

import (
	"os"
	"testing"
)

// TestNewDeck function
// Convention to start testing functions with uppercase letter
// When called, it automatically passes in values
// t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// deck should have 16 cards
	if len(d) != 16 {
		// this is a formatted string, we can pass in values
		// we need to inject the value into the string
		// we do this using %v
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// first card should be Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, got %v", d[0])
	}

	// last card should be Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubds, got %v", d[len(d)-1])
	}
}

// Testing File IO
// testing saveToFile and newDeckFromFile functions
// Long name, but if test fails, we know exactly what function
// to change without much effort
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// check for _decktesting file and remove
	os.Remove("_decktesting")

	// create new deck
	deck := newDeck()

	// save to file
	deck.saveToFile("_decktesting")

	// load from deck
	loadedDeck := newDeckFromFile("_decktesting")

	// is the deck the right length?
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	// clean up test file
	os.Remove("_decktesting")
}
