package main

import (
	"fmt"

	"github.com/vSterlin/bj/card"
)

func main() {

	d := card.NewDeck()
	for _, c := range d {
		fmt.Printf("%v\n", c)
	}
}
