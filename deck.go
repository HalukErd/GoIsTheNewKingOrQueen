package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) printDeck() {
	for i, card := range d {
		println(i, card)
	}
}

func newDeck() deck {
	numbers := []string{"Ace", "Two", "Three", "Four"}
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}

	var cards deck
	for _, suit := range suits {
		for _, number := range numbers {
			cards = append(cards, number+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, cardSize int) (deck, deck) {
	hand := d[:cardSize]
	remainingCards := d[cardSize:]
	return hand, remainingCards
}

func (d deck) toString() string {
	cardsAsStringSlice := []string(d)
	return strings.Join(cardsAsStringSlice, ",")
}

func (d deck) saveToFile(fileName string) error {
	err := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func readFromFile(filename string) deck {
	sliceByteOfCards, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	stringOfCards := string(sliceByteOfCards)

	s := strings.Split(stringOfCards, ",")
	return deck(s)
}

func (d deck) shuffleDeck() {
	//rand.Seed(time.Now().UnixNano())
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for index := range d {
		newPosition := r.Intn(len(d))
		println(newPosition)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
