package utils_test

import (
	"blackjack/deck"
	"blackjack/utils"
	"testing"
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
				{Suit: "Hearts", Value: "1"},
				{Suit: "Hearts", Value: "2"},
			},
			expectedValue: 3,
		},
		{
			name: "Total of 15",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "King"},
				{Suit: "Hearts", Value: "5"},
			},
			expectedValue: 15,
		},
		{
			name: "Total of 21",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "King"},
				{Suit: "Hearts", Value: "Ace"},
			},
			expectedValue: 21,
		},
		{
			name: "Total of 18",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "King"},
				{Suit: "Hearts", Value: "7"},
				{Suit: "Hearts", Value: "Ace"},
			},
			expectedValue: 18,
		},
		{
			name: "Total of 12",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "Ace"},
				{Suit: "Spades", Value: "Ace"},
			},
			expectedValue: 12,
		},
		{
			name: "Total of 18",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "Ace"},
				{Suit: "Spades", Value: "Ace"},
				{Suit: "Spades", Value: "6"},
			},
			expectedValue: 18,
		},
		{
			name: "Total of 18",
			hand: []*deck.Card{
				{Suit: "Hearts", Value: "Ace"},
				{Suit: "Spades", Value: "Ace"},
				{Suit: "Clubs", Value: "Ace"},
			},
			expectedValue: 13,
		},
	}

	for _, testCase := range handValueTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			value := utils.CalculateHandValue(testCase.hand)
			if value != testCase.expectedValue {
				t.Errorf("expected: %d, received: %d", testCase.expectedValue, value)
			}
		})
	}

}
