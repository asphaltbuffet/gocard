// Package gocard represents and manipulates playing cards and decks, for
// building card games that run in a terminal.
//
// # Cards
//
// A [Card] is a French-suited playing card: a [Rank] and a [Suit]. It is a
// plain comparable struct, so cards work with ==, as map keys, and in test
// assertions:
//
//	c := gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}
//	fmt.Println(c) // Q♣
//
// A card carries no game value. What a card is worth depends on the game, and
// often on the rest of the hand — an ace is 1 or 11 in blackjack, 1 in
// cribbage, 14 at high card — so valuation is a function the game supplies.
// See [Valuer], [PipValue] and [BlackjackValue].
//
// # Decks
//
// A [Deck] holds cards of a single type and is generic over that type, so
// shuffling, drawing and dealing work for any card family. [Standard] names
// the common case, a deck of [Card]:
//
//	d := gocard.NewDeck()
//	d.Shuffle(rand.New(rand.NewPCG(1, 1)))
//
//	ways, err := d.Deal(4, 5)
//
// A new deck is ordered, never pre-shuffled, so a game's tests can rely on a
// known starting order. Options adjust the composition: [WithJokers],
// [WithRanks], [WithSuits] and [WithCopies] — the last builds what a casino
// calls a shoe, which is an ordinary deck here rather than a separate type.
//
// Drawing and dealing are atomic. Asking for more cards than remain reports
// [ErrInsufficientCards] and removes nothing, so a short deal cannot leave
// some cards already handed out.
//
// # Display
//
// A card describes how it presents itself as a [Layout]: a size, a
// [BorderStyle], unstyled text, and a semantic [Accent] such as "this is a
// red suit". Renderers turn a Layout into terminal output, so a renderer
// written today can draw a card family invented tomorrow. See the
// gocard/render package for the renderer contract and a dependency-free plain
// renderer, and gocard/render/lipgloss for colour.
//
// This package itself has no external dependencies: [Card.String] always
// produces readable output with nothing imported.
//
// # Scope
//
// gocard models cards, the containers that hold them, and how a card
// presents itself. It does not model game rules — no trick-taking, melds,
// betting or scoring, and it never compares two cards to decide which wins —
// nor player and table state. Arranging several cards into a display is the
// game author's job; a Layout is structured so a renderer can measure it.
//
// CONTEXT.md in the repository root records the project's vocabulary, and
// docs/adr/ records why the API is shaped this way.
package gocard
