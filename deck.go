package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

// Any variable of type "deck"
// now gets access to this func
// func | receiver | name
func (d deck) print() {
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

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666) //anyone can read or write
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// log the error and a call to newDeck()
		// log the error and quit the program
		log.Println("Error", err)
		os.Exit(1)
	}

	stringDeck := strings.Split(string(bs), ",")
	return deck(stringDeck)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := range d {
		position := random.Intn(len(d) - 1)
		d[i], d[position] = d[position], d[i]
	}
}
