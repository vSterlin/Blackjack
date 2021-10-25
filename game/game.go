package game

import (
	"fmt"
	"strings"

	"github.com/vSterlin/bj/card"
	"github.com/vSterlin/bj/player"
)

const (
	HIT = iota + 1
	STAND
)

type Game struct {
	Player *player.Player
	Dealer *player.Player
	Deck   *card.Deck
	step   uint
}

func NewGame() *Game {
	d := card.NewDeck()
	g := &Game{&player.Player{}, &player.Player{}, &d, 0}

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
	g.step++
}

func (g *Game) PrintHands() {
	ds := "Dealer has: "
	if g.step != 1 {
		dc := []string{}
		for _, c := range g.Dealer.Hand {
			dc = append(dc, fmt.Sprintf("%s of %s", c.Value, c.Suit))
		}
		ds += strings.Join(dc, ", ")
	} else {
		ds += fmt.Sprintf("%s of %s, ***HIDDEN CARD***", g.Dealer.Hand[0].Value, g.Dealer.Hand[0].Suit)
	}

	ps := "Player has: "
	pc := []string{}
	for _, c := range g.Player.Hand {
		pc = append(pc, fmt.Sprintf("%s of %s", c.Value, c.Suit))
	}
	ps += strings.Join(pc, ", ")

	fmt.Println(ds)
	fmt.Println(ps)
}

func (g *Game) Play() {

	fmt.Println("Please choose: 1) HIT or 2) STAND")

	var choice int
	fmt.Scanf("%d", &choice)

	if choice == HIT {
		fmt.Println("HIT")
		g.Player.Hand = append(g.Player.Hand, g.Deck.TakeCard())
	} else if choice == STAND {
		// do stuff
		fmt.Println("STAND")

	}
}
