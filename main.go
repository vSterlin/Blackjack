package main

import (
	"github.com/vSterlin/bj/game"
)

func main() {

	g := game.NewGame()
	g.Deal()
	g.Deck.Print()
}
