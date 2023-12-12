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
