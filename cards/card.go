package cards

import "sort"

// Rank is the constant value of a card
type Rank int

// All possible ranks for a card
const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// Suit represents the suit of the card
type Suit int

// All possible suits for a card
const (
	Hearts Suit = iota
	Spades
	Clubs
	Diamonds
)

// RankSorter is used to sort a slice of cards by value
type RankSorter []Card

func (a RankSorter) Len() int           { return len(a) }
func (a RankSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RankSorter) Less(i, j int) bool { return a[i].rank < a[j].rank }

// Card is made up of a rank and a suit. Example, 7 of hearts,
// where 7 is the rank and hearts is the suit
type Card struct {
	rank Rank
	suit Suit
}

// NewCard returns a new card with the specified rank and suit
func NewCard(rank Rank, suit Suit) Card {
	card := Card{rank, suit}
	return card
}

// GetSuit returns the suit of the current card
func (card *Card) GetSuit() Suit {
	return card.suit
}

// GetRank returns the rank of the current card
func (card *Card) GetRank() Rank {
	return card.rank
}

// SortCards sorts provided cards in ascending order
func SortCards(cards []Card) {
	sort.Sort(RankSorter(cards))
}
