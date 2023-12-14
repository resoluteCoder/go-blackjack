package deck_test

import (
	"blackjack/deck"
	"slices"
	"strings"
	"testing"
)

func TestInitDeck(t *testing.T) {
	deck := deck.Deck{}
	deck.Init()
	if len(deck.Cards()) != 52 {
		t.Error("Deck was not initialized properly")
	}
}

func TestShuffleDeck(t *testing.T) {
	deck := deck.Deck{}
	deck.Init()
	deck.Shuffle()

	values := []string{"Ace", "2", "3", "4", "5"}
	count := 0
	for i := 0; i < 5; i++ {
		currentCard := deck.Cards()[i]
		t.Logf("%+v", currentCard)
		if strings.Contains(currentCard.Suit, "Hearts") && slices.Contains(values, currentCard.Value) {
			count++
		}
	}
	if count == 5 {
		t.Error("Deck was not shuffled properly")
	}
}

func TestGetTopCardFromDeck(t *testing.T) {
	deck := deck.Deck{}
	deck.Init()

	topCard := deck.TopCard()

	if topCard.Suit != "Hearts" && topCard.Value != "Ace" {
		t.Error("Did not get the top card")
	}

	if deck.Cards()[0].Suit != "Hearts" && deck.Cards()[0].Value != "2" {
		t.Error("Failed to remove top card")
	}
}

func TestConvertValueToWorth(t *testing.T) {
	handValueTestCases := []struct {
		name          string
		value         string
		expectedWorth int
	}{
		{
			name:          "Value of 3 worth of 3",
			value:         "3",
			expectedWorth: 3,
		},
		{
			name:          "Value of King worth of 10",
			value:         "King",
			expectedWorth: 10,
		},
		{
			name:          "Value of Ace worth of 11",
			value:         "Ace",
			expectedWorth: 11,
		},
	}
	for _, testCase := range handValueTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			worth := deck.ConvertValueToWorth(testCase.value)
			if worth != testCase.expectedWorth {
				t.Errorf("expected: %d, received: %d", testCase.expectedWorth, worth)
			}
		})
	}
}
