package main

import (
	"time"
	"math/rand"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]	
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

/** Receivers */
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// d is the reference of the actual card
// by convention, we always refer to the reference with 1 or 2 letters abreviation
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(filename string) error {
	content := []byte(d.toString())
	return ioutil.WriteFile(filename, content, 0666)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Swapping items
		d[i], d[newPosition] = d[newPosition], d[i]
	} 
}
