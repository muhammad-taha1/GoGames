package poker

import (
	"github.com/muhammad-taha1/games/cards"
)

// HandRanking orders the poker winning hands, from lowest to highest
type HandRanking int

const (
	// HighCard represents a hand of no pairs, flushes, straights etc
	HighCard HandRanking = iota
	// Pair represents two cards that have the same value and color
	Pair
	// TwoPair represents hand containing two pairs
	TwoPair
	// ThreeOfAKind represents a hand having 3 cards of the same rank
	ThreeOfAKind
	// Straight represents a hand having 5 cards of consecutive rank
	Straight
	// Flush represents a hand having 5 cards of same suit
	Flush
	// FullHouse represents a hand having a ThreeOfAKind and a Pair
	FullHouse
	// FourOfAKind represents a hand having 4 cards of the same rank
	FourOfAKind
	// StraightFlush represents a hand having 5 cards of same suit in consecutive order
	StraightFlush
	// RoyalFlush represents a hand of ten, jack, queen, king, ace of the same suit
	RoyalFlush
)

func DetermineHand(allCards []cards.Card) HandRanking {

	for _, card := range allCards {
		rank := card.GetRank()
	}
}

func SortAscending(allCards *[]cards.Card) {
	sort(allCards)
}
