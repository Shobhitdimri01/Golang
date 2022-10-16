package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//created a new type of deck
// which is a slice of string
type deck []string

func newDeck() deck {

	cards := deck{}
	CardsSuite := deck{"Spade", "Diamonds", "Hearts", "Clubs"}
	cardsValue := deck{"Ace", "Two ", "Three", "Four"}

	for _, suit := range CardsSuite {
		for _, value := range cardsValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i+1, c)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toByte() []byte {
	// a := []string(d)

	str := strings.Join(d, ",")
	b := []byte(str)

	return b
}

func (d deck) SaveFile(filename string) error {
	filebyte := d.toByte()
	err := ioutil.WriteFile(filename, filebyte, 0666)
	if err != nil {
		return err
	}
	fmt.Println("File saved with filename", filename)
	return nil
}

func (d deck) ReadFile(filename string) deck {
	filebyte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	str := strings.Split(string(filebyte), ",")
	strArray := deck(str)
	return strArray
}

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		//swapping the element
		d[i], d[newPos] = d[newPos], d[i]
	}

}
