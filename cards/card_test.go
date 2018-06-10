package cards

import (
	"fmt"
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

func TestRankSorter(t *testing.T) {
	cards := []Card{NewCard(Seven, Hearts), NewCard(Ace, Hearts), NewCard(Queen, Hearts), NewCard(Two, Hearts)}
	SortCards(cards)

	for _, card := range cards {
		fmt.Printf(fmt.Sprintf("val: %d, suit: %d\n", card.rank, card.suit))
	}
}
