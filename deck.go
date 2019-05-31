// part of the executable main package
package main

import "fmt"

// Create a new type ('deck'), a slice of strings
type deck []string

// This function will create and return a
// list of cards
// We need to annotate the type of variable the
// function is returning
// Doesn't need a receiver
func newDeck() deck {
	// starts out empty
	cards := deck{}
	// two slices: available suits, available card values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// create 2 loops (nested) to create list of 52 cards programmatically
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	// return newly created deck of cards
	return cards
}

// Loop through deck
// This function includes a receiver
// which allows any variable of type "deck"
// to have access to this print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
