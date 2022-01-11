package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := 16
	if len(d) != expectedLength {
		t.Errorf("The Length Of Deck is not %v. You got %v", expectedLength, len(d))
	}

	expectedFirstElement := "Ace of Spades"
	if d[0] != expectedFirstElement {
		t.Errorf("First element is not %v. Your first element is %v", expectedFirstElement, d[0])
	}

	expectedLastElement := "Four of Diamonds"
	if d[len(d)-1] != expectedLastElement {
		t.Errorf("Last element is not %v. Your first element is %v", expectedLastElement, d[len(d)-1])
	}
}
