package game

import (
	"github.com/vSterlin/bj/card"
	"github.com/vSterlin/bj/player"
)

type Game struct {
	Player *player.Player
	Dealer *player.Player
	Deck   card.Deck
}

func NewGame() *Game {
	d := card.NewDeck()
	d.Shuffle()
	g := &Game{&player.Player{}, &player.Player{}, d}

	return g
}
