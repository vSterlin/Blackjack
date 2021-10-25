package card

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []*Card

func NewDeck() Deck {
	d := Deck{}
	for v, rank := range ranks {
		for _, suit := range suits {
			// j, q, k need value 10
			if v > 9 {
				v = 9
			}
			d = append(d, &Card{rank, suit, v + 1})
		}
	}

	return d
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	maxN := len(*d)
	for i := range *d {
		n := rand.Intn(maxN)
		(*d)[i], (*d)[n] = (*d)[n], (*d)[i]
	}
}

func (d Deck) Print() {
	for _, c := range d {
		fmt.Printf("%v\n", c)
	}

}

func (d *Deck) TakeCard() *Card {
	c := (*d)[0]
	*d = append((*d)[:0], (*d)[1:]...)
	return c
}
