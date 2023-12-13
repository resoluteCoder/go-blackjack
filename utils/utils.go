package utils

// import (
// 	"blackjack/deck"
// 	"fmt"
// 	"slices"
// 	"strconv"
// )

// // check for known combos to avoid calculation
// // such as any face card plus an ace
// // or two aces = 12

// func ConvertValueToWorth(value string) int {
// 	var worth int
// 	faceCard := []string{"Jack", "Queen", "King"}
// 	if slices.Contains(faceCard, value) {
// 		worth = 10
// 	} else {
// 		num, err := strconv.Atoi(value)
// 		if err != nil {
// 			fmt.Println("error in converting string to int")
// 		}
// 		worth = num
// 	}
// 	if value == "Ace" {
// 		worth = 11
// 	}
// 	return worth
// }

// func CalculateHandValue(hand []*deck.Card) int {
// 	total := 0
// 	faceCard := []string{"Jack", "Queen", "King"}
// 	for _, card := range hand {
// 		if slices.Contains(faceCard, card.Value) {
// 			total += 10
// 		} else {
// 			num, err := strconv.Atoi(card.Value)
// 			if err != nil {
// 				fmt.Println("error in converting string to int")
// 			}
// 			total += num
// 		}
// 		if card.Value == "Ace" {
// 			total += 11
// 		}
// 	}

// 	if total > 21 {
// 		numOfAces := 0
// 		for _, card := range hand {
// 			if card.Value == "Ace" && numOfAces == 0 {
// 				total -= 10
// 				numOfAces++
// 			}
// 		}
// 	}
// 	return total
// }
