package main

func main() {
	//cards := newDeck()
	//hand, theDeck := deal(cards, 4)
	//hand.printDeck()
	//theDeck.printDeck()
	//cards.saveToFile("cards")
	cards := readFromFile("cards")
	cards.shuffleDeck()
	cards.printDeck()
}
