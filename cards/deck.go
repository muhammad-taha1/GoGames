package cards

import "math/rand"

// Deck is made of a slice cards
type Deck struct {
	cards []Card
}

// New returns a new deck of cards in order
func New() (*Deck, error) {

	var ranks = []Rank{
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Ten,
		Jack,
		Queen,
		King,
		Ace,
	}

	var suits = []Suit{
		Hearts,
		Spades,
		Clubs,
		Diamonds,
	}
	cards := make([]Card, 0)
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, NewCard(rank, suit))
		}
	}

	deck := Deck{cards}
	return &deck, nil
}

// Shuffle the generated deck of cards
func (deck *Deck) Shuffle() {
	for i := 0; i < len(deck.cards); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			deck.cards[r], deck.cards[i] = deck.cards[i], deck.cards[r]
		}
	}
}

// Deal returns the top card of the deck. It panics if deck is empty
func (deck *Deck) Deal() *Card {
	if len(deck.cards) <= 0 {
		panic("Deck out of cards!")
	}
	card := deck.cards[0]
	deck.cards = deck.cards[1:]
	return &card
}
