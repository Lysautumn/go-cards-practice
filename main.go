// executable package
package main

// will be called when we execute this file
func main() {
	// create a variable and give it an assignment
	// this is a long and specific way to assign this variable
	// var card string = "Ace of Spaces"

	// Go compiler  will infer it's a string because of the assignment,
	// so we can also assign the variable like this:
	// card := "Ace of Spaces"
	// The := is only when we are declaring a new variable
	// When reassigning, we can do `card = "Three of Hearts"`

	// When main is executed, will get return from newCard function
	// and assign it to card variable
	// card := newCard()
	// fmt.Println(card)

	// This is a slice, using deck type from deck.go
	//cards := deck{"Ace of Diamonds", newCard()}

	// Adding a new card to the cards slice
	// Doesn't modify existing slice, returns a new slice
	//cards = append(cards, "Six of Hearts")

	// cards is the result of calling newDeck function
	cards := newDeck()

	// create hands
	// since deal returns 2 decks, we need to create
	// variables to capture them
	hand, remainingCards := deal(cards, 5)

	// print out hand and remainingCards (we can do that
	// since both are of type deck)
	hand.print()
	remainingCards.print()

	// Iterate over cards slice
	// Using "deck" print method
	//cards.print()

	//fmt.Println(cards)
}

// will return value of card variable
// 'string' after parentheses tells Go compiler to expect function
// to return a string
//func newCard() string {
// returns a string
//return "Five of Hearts"
//}
