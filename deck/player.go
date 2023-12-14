package deck

type Player struct {
	hand         []*Card
	Bust         bool
	CurrentWorth int
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
