package main

import (
	"blackjack/deck"
)

func main() {
	dealer := deck.NewDealer()
	player := deck.NewPlayer()

	dealer.Setup(player)
}
