package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

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

// has a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// has parameters
// return multiple values from functions (deck, deck)
func deal(d deck, handSize int) (deck, deck) {

	// everything from start till > handSize
	// everything from handSize till end
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// turn a deck into slice of string, join them by comma
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	//turn the deck string in a slice of bytes
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
