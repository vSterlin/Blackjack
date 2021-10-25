package game

import (
	"fmt"
	"strings"

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

func (g *Game) PrintHands() {
	ds := "Dealer has: "
	dc := []string{}
	for _, c := range g.Dealer.Hand {
		dc = append(dc, fmt.Sprintf("%s of %s", c.Value, c.Suit))
	}
	ds += strings.Join(dc, ", ")

	ps := "Player has: "
	pc := []string{}
	for _, c := range g.Player.Hand {
		pc = append(pc, fmt.Sprintf("%s of %s", c.Value, c.Suit))
	}
	ps += strings.Join(pc, ", ")

	fmt.Println(ds)
	fmt.Println(ps)
}
