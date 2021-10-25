package card

import (
	"fmt"
	"math/rand"
)

type Deck []*Card

func NewDeck() Deck {
	d := Deck{}
	for _, value := range values {
		for _, suit := range suits {
			d = append(d, &Card{value, suit})
		}
	}

	return d
}

func (d Deck) Shuffle() {
	for i := range d {
		n := rand.Intn(52)
		d[i], d[n] = d[n], d[i]
	}
}

func (d Deck) Print() {
	for _, c := range d {
		fmt.Printf("%v\n", c)
	}

}
