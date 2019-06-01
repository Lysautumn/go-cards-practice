// executable package
package main

// will be called when we execute this file
func main() {
	//cards := newDeck()
	// call saveToFile function to create a file
	// locally with deck of cards
	// this file is generated in this project
	//cards.saveToFile("my_cards")

	// reads previously generated deck saved locally
	//cards := newDeckFromFile("my_cards")
	//cards.print()

	// create new deck of cards, shuffle cards, print
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
