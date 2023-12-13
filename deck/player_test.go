package deck_test

import (
	"blackjack/deck"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestCalculateHandValue(t *testing.T) {
	handValueTestCases := []struct {
		name          string
		hand          []*deck.Card
		expectedValue int
	}{
		{
			name: "Total of 3",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "1", Worth: 1},
				{Suit: "Hearts", Value: "2", Worth: 2},
			},
			expectedValue: 3,
		},
		{
			name: "Total of 15",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "King", Worth: 10},
				{Suit: "Hearts", Value: "5", Worth: 5},
			},
			expectedValue: 15,
		},
		{
			name: "Total of 21",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "King", Worth: 10},
				{Suit: "Hearts", Value: "Ace", Worth: 11},
			},
			expectedValue: 21,
		},
		{
			name: "Total of 18",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "King", Worth: 10},
				{Suit: "Hearts", Value: "7", Worth: 7},
				{Suit: "Hearts", Value: "Ace", Worth: 11},
			},
			expectedValue: 18,
		},
		{
			name: "Total of 12",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "Ace", Worth: 11},
				{Suit: "Spades", Value: "Ace", Worth: 11},
			},
			expectedValue: 12,
		},
		{
			name: "Total of 18",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "Ace", Worth: 11},
				{Suit: "Hearts", Value: "Ace", Worth: 11},
				{Suit: "Hearts", Value: "6", Worth: 6},
			},
			expectedValue: 18,
		},
		{
			name: "Total of 18",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "Ace", Worth: 11},
				{Suit: "Spades", Value: "Ace", Worth: 11},
				{Suit: "Clubs", Value: "Ace", Worth: 11},
			},
			expectedValue: 13,
		},
	}

	for _, testCase := range handValueTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			value := deck.CalculateHandWorth(testCase.hand)
			if value != testCase.expectedValue {
				t.Errorf("expected: %d, received: %d", testCase.expectedValue, value)
				spew.Dump(testCase.hand)
			}
		})
	}

}
