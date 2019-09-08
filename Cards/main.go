package main

func main() {
	cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// cards.saveToFile("My Card")
	// cardFromFile := newDeckFromFile("My Card")
	// // fmt.Println(cardFromFile)
	// cardFromFile.print()
	cards.shuffle()
	cards.print()
}
