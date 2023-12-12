package main

import (
	"blackjack/deck"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	dealer := deck.NewDealer()
	player := deck.NewPlayer()

	dealer.Setup(player)

	spew.Dump(dealer.Hand())
	spew.Dump(player.Hand())
}
