package deck

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
)

type Card struct {
	Suit  string
	Value string
	Worth int
}

type Deck struct {
	cards []*Card
}

func ConvertValueToWorth(value string) int {
	var worth int
	faceCard := []string{"Jack", "Queen", "King"}
	if slices.Contains(faceCard, value) {
		worth = 10
	} else {
		if value == "Ace" {
			worth = 11
			return worth
		}
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("error in converting string: %s to int", value)
		}
		worth = num
	}
	return worth
}

func (d *Deck) Init() {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range suits {
		for _, value := range values {
			d.cards = append(d.cards, &Card{
				Suit:  suit,
				Value: value,
				Worth: ConvertValueToWorth(value),
			})
		}
	}
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Cards() []*Card {
	return d.cards
}

func (d *Deck) TopCard() *Card {
	topCard := d.cards[0]
	d.cards = slices.Delete(d.cards, 0, 1)
	return topCard
}
