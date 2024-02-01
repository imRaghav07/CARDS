package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func print(d deck) {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	var deck []string
	cardSuits := []string{"spades", "diamonds", "Hearts"}
	cardValues := []string{"Ace", "one", "two", "three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, suit+"of"+value)
		}
	}
	return deck
}

func deal(size int, d deck) (deck, deck) {
	return d[0:size], d[size:]
}

func (d deck) toString() string {
	return strings.Join(([]string(d)), ",")
}

func (d deck) writeToFile(fileName string) {
	os.WriteFile(fileName, []byte(d.toString()), 0644)

}

func (d deck) readDeckFromFile(fileName string) deck {
	byteDeck, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteDeck), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
