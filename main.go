package main

import "fmt"

func main() {
	cards := newDeck()
	print(cards)
	dealOfCards, remainingCards := deal(5, cards)
	fmt.Println("A deal of cards: ", dealOfCards, "Remaining cards: ", remainingCards)
	cards.shuffle()
	fmt.Println("Post Shuffling ", cards)
	cards.writeToFile("deck")
	deckFromFile := cards.readDeckFromFile("deck")
	fmt.Println("Cards From File ", deckFromFile)

}
