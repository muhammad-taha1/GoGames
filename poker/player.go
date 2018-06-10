package poker

import (
	"fmt"

	"github.com/muhammad-taha1/games/cards"
)

// Player is a structure defining a poker player properties and functionalities
type Player struct {
	cardA cards.Card
	cardB cards.Card
	money int
}

// GetHandCards returns the two hand card player has
func (player *Player) GetHandCards() (cards.Card, cards.Card) {
	return player.cardA, player.cardB
}

// Bet function removes the input amount from player's money.
// Panic if amount is greater than money
func (player *Player) Bet(amount int) {
	if amount > player.money {
		panic(fmt.Sprintf("Cannot bet %d, only have %d", amount, player.money))
	} else {
		player.money -= amount
	}
}
