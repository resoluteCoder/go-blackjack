package ui

import (
	"blackjack/deck"
	"fmt"
)

func DisplayHeading(role string, value int) {
	// if role == "Dealer" {
	// 	fmt.Printf("%s\n", role)
	// 	return
	// }
	fmt.Printf("%s - Hand value: %d\n", role, value)
}

func DisplayHand(hand []*deck.Card, role string) {
	for _, card := range hand {
		// if role == "dealer" && index == 0 {
		// 	fmt.Println("[ Hidden ]")
		// 	continue
		// }
		fmt.Printf("%s of %s\n", card.Value, card.Suit)
	}
	fmt.Println()
}

func DisplayPlayerActions() {
	fmt.Println("1) Hit")
	fmt.Println("2) Stand")
}
