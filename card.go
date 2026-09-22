package gocard

// glyphUnknown is the one-line form of a card whose rank is absent.
const glyphUnknown = "??"

// glyphJoker is the one-line form of a joker, which has no suit pip.
const glyphJoker = "JK"

// Card is a single French-suited playing card.
//
// Card is a concrete type, not an interface: other card families are separate
// concrete types, so comparing a Card with a card of another family is a
// compile error rather than a silently false comparison. See
// docs/adr/0002-card-families-are-concrete-types.md.
//
// Card is comparable with == and usable as a map key. Its zero value is a
// card with no rank and no suit, which prints as "??".
//
// A Card carries no game value. What a card is worth depends on the game and
// often on the rest of the hand, so gocard leaves valuation to the caller —
// see [Valuer].
type Card struct {
	// Rank is the card's identity within its suit.
	Rank Rank

	// Suit is the card's symbol group. It is [NoSuit] for a joker.
	Suit Suit
}

// String returns the card's one-line plain-text form: a rank symbol followed
// by a suit pip, such as "Q♣".
//
// A joker is "JK" and a card with no rank is "??". String needs no renderer
// and no external dependencies, so [fmt.Println] on a Card always produces
// readable output.
func (c Card) String() string {
	if c.IsJoker() {
		return glyphJoker
	}

	symbol := c.rankSymbol()
	if symbol == "" {
		return glyphUnknown
	}

	return symbol + c.Suit.Symbol()
}

// rankSymbol returns the short display form of the card's rank, or the empty
// string if the rank is absent or unrecognised.
func (c Card) rankSymbol() string {
	switch c.Rank {
	case Ace:
		return "A"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Joker, NoRank:
		return ""
	}

	return ""
}

// IsJoker reports whether c is a joker.
func (c Card) IsJoker() bool {
	return c.Rank == Joker
}

// Valid reports whether c is a card that could appear in a deck: a joker, or
// a card with both a rank and a suit.
//
// A joker is checked by rank alone, so a card with rank [Joker] and a suit
// reports valid even though a joker conventionally has no suit. Such a card
// prints and renders as a plain joker, its suit ignored.
func (c Card) Valid() bool {
	if c.IsJoker() {
		return true
	}

	return c.Rank.Valid() && c.Suit.Valid()
}
