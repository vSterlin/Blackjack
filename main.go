package main

import (
	"github.com/vSterlin/bj/card"
)

func main() {

	d := card.NewDeck()
	d.Shuffle()
	d.Print()
}
