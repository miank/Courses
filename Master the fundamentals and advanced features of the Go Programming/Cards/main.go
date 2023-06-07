package main

func main() {
	// cards := newDeck()

	// hand, remaniningDeck := deal(cards, 5)
	// hand.print()
	// remaniningDeck.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards") // Saving the content to a file
	// // Load the data from a file.

	// fmt.Println(cards.toString())

	cards := newDeckFromFile("../../my_cards")
	cards.print()

}
