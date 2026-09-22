package gocard

// minCopies is the fewest copies a deck can be built from.
const minCopies = 1

// WithJokers adds n jokers to each copy of the deck.
//
// Jokers are opt-in: a standard deck has 52 cards, and
// NewDeck(WithJokers(2)) has 54. A joker is a [Card] with rank [Joker] and no
// suit, so a deck containing jokers still holds a single card type.
//
// A negative n adds no jokers.
func WithJokers(n int) Option {
	return func(cfg *config) {
		if n < 0 {
			n = 0
		}

		cfg.jokers = n
	}
}

// WithRanks restricts the deck to the given ranks, for games that do not use
// the full sequence — euchre uses nine through ace, for example.
//
// Ranks appear in the order supplied. Passing no ranks leaves the standard
// ace-through-king sequence in place.
func WithRanks(ranks ...Rank) Option {
	return func(cfg *config) {
		if len(ranks) == 0 {
			return
		}

		cfg.ranks = ranks
	}
}

// WithSuits restricts the deck to the given suits.
//
// Suits appear in the order supplied. Passing no suits leaves the standard
// four in place.
func WithSuits(suits ...Suit) Option {
	return func(cfg *config) {
		if len(suits) == 0 {
			return
		}

		cfg.suits = suits
	}
}

// WithCopies builds the deck from n copies of its composition.
//
// This is how to build the multi-deck arrangement a casino calls a shoe:
// NewDeck(WithCopies(6)) holds 312 cards. gocard has no separate shoe type,
// because a shoe draws and shuffles exactly like any other deck.
//
// An n below one builds a single copy.
func WithCopies(n int) Option {
	return func(cfg *config) {
		if n < minCopies {
			n = minCopies
		}

		cfg.copies = n
	}
}
