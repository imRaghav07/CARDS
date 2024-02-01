package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 12 {
		t.Errorf("Tets case failed. Length of the deck do not match with expected length. Expected Length: %v Result length: %v", 12, len(cards))
	}
}

func TestWriteToFileAndReadDeckFromFile(t *testing.T) {

	os.Remove("TestDeckFile")
	cards := newDeck()
	cards.writeToFile("TestDeckFile")
	d := cards.readDeckFromFile("TestDeckFile")
	if len(d) != 12 {
		t.Errorf("Tets case failed. Length of the deck do not match with expected length. Expected Length: %v Result length: %v", 12, len(cards))
	}
	os.Remove("TestDeckFile")
}
