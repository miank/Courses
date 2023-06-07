package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type 'deck'
// which is a slice of strings

// Consider this is as a class decleration in terms of OO Language
type deck []string // Borrows the behaviouur of slice of string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{
		"Spades", "Hearts", "Diamonds", "Clubs",
	}

	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
	}

	for _, suits := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suits)
		}
	}

	return cards

}

func (d deck) print() { // d is a reference to the cards variable.
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error :", err)
	}
	os.Exit(1)
	s := strings.Split(string(bs), ",")

	return  deck(s)

}
