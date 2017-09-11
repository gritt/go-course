package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
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

// read bytes from file and return as a new deck
func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// #1 - log the err and return a call to newDeck()
		// #2 - log the err and exit
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// turn bs slice os strings
	s := strings.Split(string(bs), ",")

	// cause the deck type is a slice os string
	return deck(s)
}

// randomize the working deck, without return
func (d deck)shuffle() {

	// current time unique unixnano, is the source for rand numbers
	s := rand.NewSource(time.Now().UnixNano())

	r := rand.New(s)

	for i := range d {
		np := r.Intn(len(d) - 1)

		d[i], d[np] = d[np], d[i]
	}
}
