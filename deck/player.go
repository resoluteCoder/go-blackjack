package deck

type Player struct {
	hand         []*Card
	Bust         bool
	CurrentWorth int
}

func CalculateHandWorth(hand []*Card) int {
	total := 0
	for _, card := range hand {
		total += card.Worth
	}
	if total > 21 {
		for _, card := range hand {
			if card.Value == "Ace" {
				card.Worth = 1
				total = 0
				for _, card := range hand {
					total += card.Worth
				}
				if total > 21 {
					continue
				} else {
					break
				}
			}
		}
	}
	return total
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Hit(c *Card) {
	p.hand = append(p.hand, c)
}

func (p *Player) Hand() []*Card {
	return p.hand
}
