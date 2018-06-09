package cards

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
