package game

import (
	"github.com/vSterlin/bj/card"
	"github.com/vSterlin/bj/player"
)

type Game struct {
	Player *player.Player
	Dealer *player.Player
	Deck   *card.Deck
}

func NewGame() *Game {
	d := card.NewDeck()
	g := &Game{&player.Player{}, &player.Player{}, &d}

	return g
}

func (g *Game) Deal() {
	g.Deck.Shuffle()
	// card to player, card to dealer, twice
	for i := 0; i < 2; i++ {
		c := g.Deck.TakeCard()
		g.Player.Hand = append(g.Player.Hand, c)
		c = g.Deck.TakeCard()
		g.Dealer.Hand = append(g.Dealer.Hand, c)
	}
}
