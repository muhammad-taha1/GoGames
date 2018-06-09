package cards

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	card := NewCard(Seven, Hearts)

	if card.suit != Hearts {
		t.Errorf("invalid suit. got: %d, want: %d", card.suit, Hearts)
	}

	if card.rank != Seven {
		t.Errorf("invalid rank. got: %d, want: %d", card.rank, Seven)
	}

}
