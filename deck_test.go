package main

import "testing"

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
