package card

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
