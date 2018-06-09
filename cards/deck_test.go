package cards

import (
	"testing"
)

func TestNew(t *testing.T) {
	deck, _ := New()

	if len(deck.cards) != 52 {
		t.Errorf("Deck doesnt have 52 cards! got: %d, want: %d", len(deck.cards), 52)
	}
}

func TestShuffle(t *testing.T) {
	deck, _ := New()

	deck.Shuffle()

	if deck.cards[0].rank == 2 && deck.cards[0].suit == 0 && deck.cards[1].rank == 3 && deck.cards[1].suit == 0 {
		t.Errorf("Deck is not properly shuffled")
	}
}

func TestDeal(t *testing.T) {
	deck, _ := New()
	card := deck.Deal()

	if card.suit != 0 && card.rank != 2 {
		t.Errorf("incorrect card dealt")
	}

	if len(deck.cards) != 51 {
		t.Errorf("card is not removed from deck")
	}
}
