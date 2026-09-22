package gocard

// blackjackFaceValue is what a jack, queen or king counts in blackjack.
const blackjackFaceValue = 10

// blackjackAceHigh is the high value of an ace in blackjack. A caller
// holding a hand that would bust must count the ace as 1 instead.
const blackjackAceHigh = 11

// Valuer maps a card to a number for one particular game.
//
// gocard stores no value on a card, because a card's worth is a property of
// the card in a game rather than of the card itself: an ace is 1 or 11 in
// blackjack depending on the rest of the hand, 1 in cribbage and 14 at high
// card. Games supply a Valuer instead, and this package ships a few common
// conventions such as [PipValue] and [BlackjackValue].
type Valuer[T any] func(T) int

// PipValue returns c counted with ace low and the court cards continuing the
// sequence: ace 1, jack 11, queen 12, king 13.
//
// This is [Rank.Ordinal] expressed as a [Valuer]. Cards with no position in
// that sequence, such as a joker, count 0.
func PipValue(c Card) int {
	return c.Rank.Ordinal()
}

// BlackjackValue returns c counted as in blackjack: court cards are 10 and an
// ace is 11.
//
// An ace is worth 1 or 11 depending on the whole hand, which no per-card
// function can know. BlackjackValue always returns the high value; a caller
// totalling a hand must demote aces to 1 while the total would otherwise
// exceed 21. A joker counts 0.
func BlackjackValue(c Card) int {
	switch c.Rank {
	case Ace:
		return blackjackAceHigh
	case Jack, Queen, King:
		return blackjackFaceValue
	case Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten:
		return c.Rank.Ordinal()
	case NoRank, Joker:
		return 0
	}

	return 0
}
