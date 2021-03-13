package cards

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck extends type string
type Deck []string

var suits [4]string = [4]string{
	"Spades", "Clubs", "Diamonds", "Hearts",
}
var values [13]string = [13]string{
	"Ace", "Two", "Three", "Four", "Five", "Six",
	"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King",
}

// NewDeck creates a whole deck with 52 cards
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

// NewDeckFromFile reads a file named "filename" and
// returns a Deck and an error
func NewDeckFromFile(filename string) (Deck, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Deck(strings.Split(string(data), ",")), nil
}

// Deal can be used to deal cards and returns
// the dealt hand and the remaining deck
func Deal(d Deck, handSize int) (Deck, Deck) {
	if handSize > len(d) {
		handSize = len(d)
	}

	return d[handSize:], d[:handSize]
}

// ToString convert your current deck to a string
func (d Deck) ToString() string {
	return strings.Join(d, ",")
}

// SaveToFile writes you deck into a file named filename
func (d Deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.ToString()), 0666)
}

func (d Deck) Shuffle() error {
	if len(d) == 0 {
		return errors.New("Deck is empty")
	}

	randGenerator := rand.New(rand.NewSource(time.Now().Unix()))

	for i := range d {
		n := randGenerator.Intn(len(d))
		d[i], d[n] = d[n], d[i]
	}

	return nil
}

// Print prints the deck into the standard output
func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
