package main

import "fmt"

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// has to replace vars which are not used by _
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// return multiple values from functions (deck, deck)
func deal(d deck, handSize int) (deck, deck) {

	// everything from start till > handSize
	// everything from handSize till end
	return d[:handSize], d[handSize:]
}
