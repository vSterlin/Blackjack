package card

type Card struct {
	Rank  string
	Suit  string
	Value int
}

var ranks = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var suits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}
