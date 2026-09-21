package gocard

//go:generate stringer -type=Rank -trimprefix=

// Rank is the identity of a [Card] within its suit.
//
// Rank is not a strength ordering: gocard never decides which rank beats
// which, because that is a game rule. Use [Rank.Ordinal] for the common
// Ace-low numbering, or supply a game-specific valuation.
//
// The zero value, [NoRank], means the rank is absent rather than Ace.
type Rank uint8

// Ranks of a French-suited deck. NoRank is the zero value and means absent.
//
// Joker is a Rank rather than a separate card type, so a deck containing
// jokers still holds a single card type.
const (
	NoRank Rank = iota
	Ace
	Two
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
	Joker
)

// ordinalKing is the Ace-low ordinal of [King], the highest ranked card
// that has an ordinal.
//
//nolint:unused // documents the maximum ordinal value
const ordinalKing = 13

// Ordinal returns one conventional numbering of r: Ace is 1 through King is
// 13.
//
// This is a convenience for games that need a simple number, not an
// authority on which card wins. Ranks without a position in that sequence —
// [NoRank] and [Joker] — return 0.
func (r Rank) Ordinal() int {
	switch r {
	case Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King:
		return int(r)
	case NoRank, Joker:
		return 0
	}

	return 0
}

// Valid reports whether r is a rank this package defines. [NoRank] is not
// valid, because it represents the absence of a rank.
func (r Rank) Valid() bool {
	return r >= Ace && r <= Joker
}
