package cards

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	newDeck := NewDeck()

	t.Run("test returned deck length", func(t *testing.T) {
		expected := 52
		got := len(newDeck)

		if got != expected {
			t.Errorf("Expected length %d, got %d", expected, got)
		}
	})

	t.Run("test first card", func(t *testing.T) {
		expected := "Ace of Spades"
		got := newDeck[0]

		if got != expected {
			t.Errorf("Expected first card to be %s, got %s", expected, got)
		}
	})

	t.Run("test last card", func(t *testing.T) {
		expected := "King of Hearts"
		got := newDeck[len(newDeck)-1]

		if got != expected {
			t.Errorf("Expected last card to be %s, got %s", expected, got)
		}
	})
}
