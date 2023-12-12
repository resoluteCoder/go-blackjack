package deck

type Dealer struct {
	deck *Deck
	hand []*Card
}

func NewDealer() *Dealer {
	deck := &Deck{}
	deck.Init()
	deck.Shuffle()

	return &Dealer{
		deck: deck,
	}
}

func (d *Dealer) Setup(p *Player) {
	d.Hit()
	p.Hit(d.DealCard())
	d.Hit()
	p.Hit(d.DealCard())
}

func (d *Dealer) DealCard() *Card {
	return d.deck.TopCard()
}

func (d *Dealer) Hit() {
	d.hand = append(d.hand, d.deck.TopCard())
}

func (d *Dealer) Hand() []*Card {
	return d.hand
}
