package main

import "fmt"

type deck []string

func newDeck() deck { //returns a value that is type string. deck equals string
	cards := deck{} //created a varible using " :="

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { //usually you would use "i" for the index in a for loop but since we aren't actually using it anywhere else we ca
		// replace it with an "_" as this tells Go that we know there is a varible there but we dont want to use it
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

func deal(d deck, handSize int) (deck, deck) { //the (deck, deck) is telling Go that we would like to return two values that are the same type
	return d[:handSize], d[handSize:] //it is important to state the two types of values we want to get as Go needs to see that we want 2 values
}
