package main

import (
	"blackjack/deck"
	"blackjack/ui"
	"fmt"
	"time"
)

func main() {
	// create dealer
	// create player
	dealer := deck.NewDealer()
	player := deck.NewPlayer()

	// setup dealer and player
	dealer.Setup(player)
	// display dealer and player cards

	for {
		ui.DisplayHeading("Dealer", dealer.CurrentWorth)
		ui.DisplayHand(dealer.Hand(), "dealer")

		ui.DisplayHeading("Player", player.CurrentWorth)
		ui.DisplayHand(player.Hand(), "player")
		ui.DisplayPlayerActions()

		var input string
		fmt.Scan(&input)

		if input == "1" {
			player.Hit(dealer.DealCard())
			player.CurrentWorth = deck.CalculateHandWorth(player.Hand())
			if player.CurrentWorth > 21 {
				player.Bust = true
				fmt.Println("The player Busts!")
				fmt.Println("The dealer wins!")
				break
			}
		}

		if input == "2" {
			break
		}
	}
	ui.DisplayHeading("Dealer", dealer.CurrentWorth)
	ui.DisplayHand(dealer.Hand(), "dealer")

	ui.DisplayHeading("Player", player.CurrentWorth)
	ui.DisplayHand(player.Hand(), "player")
	// repeat until stand
	// have dealer hit until 21, bust, or beats the player
	// finish
	if !player.Bust {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("The dealer hits!")
			dealer.Hit()
			dealer.CurrentWorth = deck.CalculateHandWorth(dealer.Hand())
			if dealer.CurrentWorth > 21 {
				fmt.Println("The dealer Busts!")
				fmt.Println("The player wins!")
				break
			}
			if dealer.CurrentWorth > player.CurrentWorth || dealer.CurrentWorth == 21 {
				fmt.Println("The dealer wins!")
				break
			}
			ui.DisplayHeading("Dealer", dealer.CurrentWorth)
			ui.DisplayHand(dealer.Hand(), "dealer")
		}
		ui.DisplayHeading("Dealer", dealer.CurrentWorth)
		ui.DisplayHand(dealer.Hand(), "dealer")
		ui.DisplayHeading("Player", player.CurrentWorth)
		ui.DisplayHand(player.Hand(), "player")
	}
}
