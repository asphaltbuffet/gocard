package gocard

// standardRanks are the ranks of a French-suited deck in ascending order,
// excluding [Joker], which is opt-in.
var standardRanks = []Rank{
	Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King,
}

// standardSuits are the suits of a French-suited deck in ascending order.
var standardSuits = []Suit{Clubs, Diamonds, Hearts, Spades}

// config holds the settings an [Option] adjusts before a deck is built.
type config struct {
	ranks  []Rank
	suits  []Suit
	jokers int
	copies int
}

// newConfig returns the settings for an unmodified standard deck.
func newConfig() config {
	return config{
		ranks:  standardRanks,
		suits:  standardSuits,
		jokers: 0,
		copies: 1,
	}
}

// Option adjusts how [NewDeck] builds a deck.
//
// Option is opaque: use the provided constructors such as [WithJokers],
// [WithRanks] and [WithCopies].
type Option func(*config)

// Deck is an ordered collection of cards of a single card type, from which
// cards are drawn.
//
// Deck is generic over its card type so that shuffling, drawing and dealing
// are written once and work for any card family, including card types
// defined outside this package. Because a deck holds one card type, mixing
// families is a compile error rather than a runtime check. See
// docs/adr/0002-card-families-are-concrete-types.md.
//
// A Deck is held and passed as a pointer, and its methods mutate it in
// place. Copying a Deck value shares the underlying cards; use [Deck.Clone]
// for an independent copy.
//
// The zero value is an empty deck, ready to be added to or drawn from — a
// draw from it reports [ErrInsufficientCards].
type Deck[T any] struct {
	cards []T
}

// Standard is a deck of French-suited playing cards, the common case.
type Standard = Deck[Card]

// NewDeck returns a standard 52-card deck of French-suited playing cards.
//
// The deck is ordered, never pre-shuffled: cards run ace through king in
// clubs, then diamonds, hearts and spades. Deterministic construction is what
// lets a game's tests rely on a known starting order — call [Deck.Shuffle]
// explicitly to randomise.
//
// Options adjust the composition: [WithJokers] adds jokers, [WithRanks]
// restricts the ranks for games such as euchre, and [WithCopies] builds the
// multi-deck arrangement a casino calls a shoe.
func NewDeck(opts ...Option) *Standard {
	cfg := newConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	cards := make([]Card, 0, len(cfg.suits)*len(cfg.ranks)*cfg.copies+cfg.jokers*cfg.copies)

	for range cfg.copies {
		for _, suit := range cfg.suits {
			for _, rank := range cfg.ranks {
				cards = append(cards, Card{Rank: rank, Suit: suit})
			}
		}

		for range cfg.jokers {
			cards = append(cards, Card{Rank: Joker})
		}
	}

	return &Deck[Card]{cards: cards}
}

// NewDeckOf returns a deck holding the given cards, in the order supplied.
//
// Use it to build a deck of a card type this package does not define, or to
// set up a known deck in a test.
func NewDeckOf[T any](cards ...T) *Deck[T] {
	return &Deck[T]{cards: append([]T(nil), cards...)}
}

// Len returns the number of cards remaining in the deck.
func (d *Deck[T]) Len() int {
	return len(d.cards)
}

// Cards returns the remaining cards in order, top first.
//
// The result is a copy: modifying it does not affect the deck.
func (d *Deck[T]) Cards() []T {
	return append([]T(nil), d.cards...)
}

// Clone returns an independent copy of the deck. Drawing from the copy does
// not affect the original.
func (d *Deck[T]) Clone() *Deck[T] {
	return &Deck[T]{cards: d.Cards()}
}
