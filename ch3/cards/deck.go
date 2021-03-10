package cards

import "fmt"

// Deck extends type string
type Deck []string

var suits [4]string = [4]string{
	"Spades", "Clubs", "Diamonds", "Hearts",
}
var values [13]string = [13]string{
	"Ace", "Two", "Three", "Four", "Five", "Six",
	"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King",
}

func NewDeck() Deck {
	cards := Deck{}

	for _, suit := range suits {
		for _, value := range values {
			card := fmt.Sprintf("%s of %s", value, suit)
			cards = append(cards, card)
		}
	}

	return cards
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	if handSize > len(d) {
		handSize = len(d)
	}

	return d[handSize:], d[:handSize]
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
