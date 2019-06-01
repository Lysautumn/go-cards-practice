// part of the executable main package
package main

// multiple imports
import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

// function to deal a hand of cards, returning hand and remaining cards (decks)
// function contains argument: pass in a deck
// and an integer representing the size of the hand
// 'd' and 'handSize' is the variable name we will use
// within the function (d is convention)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// helper function takes a deck to convert to string
// receiver function
func (d deck) toString() string {
	// I want to convert my deck into a slice of strings
	// deck is a slice of strings (we extended the string type
	// when we created the deck type) so this conversion is
	// straightforward
	//[]string(d)

	// join values in slice into comma-separated values in a string
	// built-in package "strings", use Join function
	return strings.Join([]string(d), ",")
}

// function to save byte slice to local machine, returns
// error if one is produced
// receiver function
// 0666 permissions means anyone can read and write this file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
