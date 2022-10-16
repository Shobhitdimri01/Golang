package main

import (
	"os"
	"testing"
)



func TestNewDeck(t *testing.T) {
	d := newDeck()
	var expected_len int = 16
	if len(d) != expected_len{
		t.Errorf("Expected deck len of %v but got %v",expected_len,len(d))
	}
}

func TestCard(t *testing.T){
	d := newDeck()
	firstCard := "Ace of Spade"
	if d[0] != firstCard {
		t.Errorf("First Card expected %v ,but got %v ",firstCard,d[0])
	}

	lastCard :="Four of Clubs"
	if d[len(d)-1] != lastCard {
		t.Errorf("Last Card expected %v ,but got %v ",lastCard,d[15])
	}
}

func TestSaveFile(t *testing.T){
	os.Remove("_decktesting")

	d := newDeck()
	d.SaveFile("_decktesting")
	loadedDeck :=d.ReadFile("_decktesting")

	length := 13
	if len(loadedDeck) != length {
		t.Errorf("Expected len %v got len %v",length,len(loadedDeck))
	}

	os.Remove("_decktesting")
}