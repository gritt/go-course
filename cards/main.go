package main

func main() {

	cards := newDeck()

	// for multiple return values (both are deck)
	hand, remainingDeck := deal(cards, 5)

	hand.print()

	remainingDeck.print()
}
