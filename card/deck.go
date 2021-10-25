package card

import (
	"fmt"
	"math/rand"
	"time"
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

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := range *d {
		n := rand.Intn(52)
		fmt.Println(n)
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
