package main

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("my_cards")

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards.print()
}
