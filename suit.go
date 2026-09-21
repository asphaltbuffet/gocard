package gocard

//go:generate stringer -type=Suit -trimprefix=

// Suit is the symbol group a [Card] belongs to.
//
// A card may have no suit: a joker is a card whose suit is absent. The zero
// value, [NoSuit], means absent rather than Clubs.
type Suit uint8

// Suits of a French-suited deck, in the conventional ascending order used by
// bridge. NoSuit is the zero value and means absent.
const (
	NoSuit Suit = iota
	Clubs
	Diamonds
	Hearts
	Spades
)

// Symbol returns the Unicode pip for s, or the empty string for [NoSuit].
func (s Suit) Symbol() string {
	switch s {
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	case Spades:
		return "♠"
	case NoSuit:
		return ""
	}

	return ""
}

// Valid reports whether s is a suit this package defines. [NoSuit] is not
// valid, because it represents the absence of a suit.
func (s Suit) Valid() bool {
	return s >= Clubs && s <= Spades
}
