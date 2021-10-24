package card

import (
	"fmt"
	"math/rand"
)

type deck []*card

func NewDeck() deck {
	d := deck{}
	for _, value := range values {
		for _, suit := range suits {
			d = append(d, &card{value, suit})
		}
	}

	return d
}

func (d deck) Shuffle() {
	for i := range d {
		n := rand.Intn(52)
		d[i], d[n] = d[n], d[i]
	}
}

func (d deck) Print() {
	for _, c := range d {
		fmt.Printf("%v\n", c)
	}
}
