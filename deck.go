// part of the executable main package
package main

// multiple imports
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// function reads local my_cards file
// return a deck
func newDeckFromFile(filename string) deck {
	// creating variables to represent info
	// returned from ReadFile function: the byte slice
	// and an error object
	bs, err := ioutil.ReadFile(filename)
	// if there is an error, error handling
	// error handling is tough because it depends on the error
	if err != nil {
		// If something goes wrong, what do I want to happen?
		// Option 1 - log error and return call to newDeck
		// Option 2 - log error and entirely quit the program because
		// something is really wrong
		fmt.Println("Error:", err)
		// Use Go os package to quit the program
		os.Exit(1)
	}
	// we have a value in our byte slice we can turn into a deck
	// convert byte slice into a string using string(bs)
	// convert string into slice of string
	// use strings package again, Split function
	// assign return to "s" variable
	s := strings.Split(string(bs), ",")

	// type conversion to change slice of string into deck
	// we can do this because deck type extends slice of string
	// return it
	return deck(s)
}

// function to shuffle
// take the deck, randomize the order inside of it
func (d deck) shuffle() {
	// make new truly random generator
	// create a new source
	// pass in int64, we want it to be a little different every time
	// use time now function, UnixNano, this will make it random
	source := rand.NewSource(time.Now().UnixNano())

	// create new Rand value to use with Intn
	r := rand.New(source)

	for i := range d {
		// create random number between 0 and length of array - 1
		newPosition := r.Intn(len(d) - 1)

		// swap cards at each position
		// this works (sort of), but last 4 are always the same
		// the rand.Intn is a pseudo random generator, it always uses
		// the same seed, so we always get the same sequence
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
