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
	LOST
)

const (
	PLAYER = iota
	DEALER
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
			dc = append(dc, fmt.Sprintf("%s of %s", c.Rank, c.Suit))
		}
		ds += strings.Join(dc, ", ")
	} else {
		ds += fmt.Sprintf("%s of %s, ***HIDDEN CARD***", g.Dealer.Hand[0].Rank, g.Dealer.Hand[0].Suit)
	}

	ps := "Player has: "
	pc := []string{}
	for _, c := range g.Player.Hand {
		pc = append(pc, fmt.Sprintf("%s of %s", c.Rank, c.Suit))
	}
	ps += strings.Join(pc, ", ")

	fmt.Println(ds)
	fmt.Println(ps)
}

func (g *Game) Play() int {
	g.step++

	fmt.Println("Please choose: 1) HIT or 2) STAND")

	var choice int
	fmt.Scanf("%d", &choice)

	if choice == HIT {
		fmt.Println("HIT")
		g.Player.Hand = append(g.Player.Hand, g.Deck.TakeCard())
		pt := g.GetTotal(PLAYER)
		if pt > 21 {
			fmt.Println("LOST")
			return LOST
		}

	} else if choice == STAND {
		// let dealer take card if necessary
		// and determine who wins
		fmt.Println("STAND")
		return STAND
	}
	// temporary
	return STAND
}

func (g *Game) DealerPlay() {
	g.step++

	dt := g.GetTotal(DEALER)

	for dt < 17 {
		g.Dealer.Hand = append(g.Dealer.Hand, g.Deck.TakeCard())
		dt = g.GetTotal(DEALER)
	}

}

func (g *Game) GetTotal(c int) int {

	// chooses whose total to check
	var p *player.Player

	if c == PLAYER {
		p = g.Player
	} else {
		p = g.Dealer
	}

	aceCount := 0
	sum := 0
	for _, c := range p.Hand {
		if c.Value == 1 {
			aceCount++
		}
		sum += c.Value
	}

	// for length of aceCount add 10
	// until reached over 21
	for i := 0; i < aceCount; i++ {
		if sum+10 <= 21 {
			sum += 10
		}
	}

	return sum
}

func (g *Game) DetermineWinner() string {
	pt := g.GetTotal(PLAYER)
	dt := g.GetTotal(DEALER)

	if pt > dt {
		res := "YOU HAVE WON!!!"
		if pt == 21 {
			res += " BLACKJACK!!!"
		}
		return res
	}
	return "YOU HAVE LOST!!!"
}

func (g *Game) Cleanup() {
	g.Player.Hand = g.Player.Hand[:0]
	g.Dealer.Hand = g.Dealer.Hand[:0]
}

func (g *Game) Loop() {

	g.Deal()
	for {
		g.PrintHands()
		res := g.Play()
		if res == LOST {
			g.Cleanup()
			fmt.Println("YOU HAVE LOST!!!")
			break
		}
		g.PrintHands()
		g.DealerPlay()
		s := g.DetermineWinner()
		g.Cleanup()
		fmt.Println(s)
	}
}
