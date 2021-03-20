package cards

import (
	"os"
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

func TestSaveToFile(t *testing.T) {
	tempDir := os.TempDir()
	f, err := os.CreateTemp(tempDir, "deck")
	if err != nil {
		t.Errorf("An error occured while creating temp file: %s", err)
	}
	deck := NewDeck()
	err = deck.SaveToFile(f.Name())

	defer os.Remove(f.Name())

	t.Run("test no error occurs after saving file", func(t *testing.T) {
		if err != nil {
			t.Errorf("Expected not error, got %s", err)
		}
	})

	t.Run("test saved file is not empty", func(t *testing.T) {
		stat, _ := os.Stat(f.Name())
		if stat.Size() == 0 {
			t.Errorf("Expected file to have size > 0, got 0")
		}
	})
}
