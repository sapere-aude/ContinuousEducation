package main

import (
	"fmt"
)

func main() {
	cards := newDeckFromFile("my_cards")
	fmt.Println(cards)
	//hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
}
