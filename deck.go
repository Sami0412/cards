package main

import "fmt"

//Create a new type of 'deck'
//which is a slice of strings
type deck []string

//Function that creates a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//Receiver function that prints the deck contents
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//(deck, deck) is the return type - we expect two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	[]string(d)
}