# Context: gocard

The domain glossary for this repo. Skills read this to learn the project's
language before proposing changes; see `docs/agents/domain.md` for the rules.

This file is **lazily filled**. Add a term when a real decision forces you to
pin its meaning down — not speculatively. An empty section is honest; a
guessed definition is worse than none.

## Domain

gocard models **cards, the containers that hold them, and how a card presents
itself on a terminal**. It is game-agnostic: it knows a Deck can be shuffled and
drawn from, and that a Card can describe how it should be drawn, but it knows
nothing about whose turn it is or who won.

Deliberately **not** modelled:

- **Game rules** — no trick-taking, melds, betting, scoring or legal-move
  validation. A Card's value is data the game interprets; gocard never compares
  two cards to decide which one wins.
- **Player and table state** — no Player, Seat, Pot or Round.
- **Multi-card composition** — arranging several cards into a fanned hand or a
  table layout is the game author's responsibility for now. Layout is designed
  so this can be added later without a breaking change.

## Glossary

Terms below are the vocabulary the exported API should use. Once a term is
exported it is part of the semver contract, so settle the name here first.

| Term | Definition | Notes |
| --- | --- | --- |
| **Card** | A single French-suited playing card. A concrete type, not an interface — a Card is *this* kind of card. Other card families (Tarot, Magic) are separate concrete types, so equality between families is a compile error rather than a silent `false`. | Comparable with `==`; usable as a map key |
| **Rank** | The identity of a Card within its suit. Not a strength ordering — gocard never decides which rank beats which. `Joker` is a Rank. | Zero value means *absent*, not Ace |
| **Ordinal** | One conventional numbering of a Rank (Ace 1 … King 13). *A* convenient ordering, not *the* ordering — offered so simple games work out of the box, never as an authority on which card wins. | |
| **Valuation** | A game-supplied function mapping a Card to a number. Value is a property of a card *in a game*, never of the card itself: an Ace is 1 or 11 in blackjack depending on the rest of the hand, 1 in cribbage, 14 at high card. gocard therefore has **no `Value` field** — it ships named valuation functions for common conventions and lets games supply their own. | |
| **Suit** | The symbol group a Card belongs to. A Card may have **no suit** — a Joker is a Card whose suit is absent. | Zero value means *absent*, not Clubs |
| **Joker** | A Card with rank `Joker` and no suit. A member of the Card family, not a separate type, so a Deck containing jokers stays homogeneous. | Opt-in: a standard Deck has 52, not 54 |
| **Deck** | An ordered collection of cards **of a single card type**, from which cards are drawn. Generic over that type, so the shuffle/draw machinery is written once and reused by every card family, including third-party ones. Held and passed as a pointer; methods mutate it. | `Deck[T]`; `Standard` aliases `Deck[Card]` |
| **Top** | The end of a Deck that cards are drawn from: index 0. "The deck in order" therefore reads top-first, so test fixtures and printed output match intuition. | |
| **Shoe** | Not a type. A multi-deck shoe is a Deck built from N copies of a standard one — `NewDeck(WithCopies(6))`. Nothing about drawing or shuffling differs, so a separate type would only duplicate the Deck. | Avoid the word in the API |
| **Ordered** | The state of a freshly constructed Deck: a defined, documented sequence — never pre-shuffled. Shuffling is one explicit call, and an unshuffled deck is what makes tests deterministic. | |
| **Draw** | Remove one or more cards from the top of a Deck and return them. **Mutates** the Deck. Atomic: drawing more cards than remain returns `ErrInsufficientCards` and removes nothing, so a short deal fails cleanly rather than leaving some players holding cards. | Never panics; never partially succeeds |
| **Deal** | Distribute cards from the top of a Deck into several ways at once, one card to each way in turn — the order a dealer uses, which differs from taking each way's cards in succession. Atomic, like Draw. | `Deal(ways, cardsPerWay int)` |
| **Way** | One of the separate groupings a Deal distributes into: "a four-way deal". A way is just a slice of cards — gocard has no notion of who holds one, which is why this word is used instead of *hand*. | Avoids the player/table boundary |
| **Shuffle** | Reorder a Deck in place using a caller-supplied random source, so games are deterministically testable. A zero-argument convenience form uses a default source. | |
| **Layout** | A card's own description of *how it presents itself*: size, border style, and content, as a single bordered box. Author-controlled, so a taller Tarot card needs no change to any renderer. Expressed in gocard's own types — never a third-party library's — so the core stays dependency-free. | |
| **Accent** | A *semantic* colour role on a Layout (e.g. "this is a red suit"), not a literal colour. The renderer's theme resolves it, which is what lets a plain renderer ignore colour entirely and a colour renderer be re-themed without touching cards. | |
| **Renderer** | Turns a Layout into terminal output for one output mode. Defined in `gocard/render`, not the core — the consumer owns the interface, so third parties can add renderers without gocard exporting anything for them. Knows nothing about card types. | |
| **Glyph** | A Card's one-line plain-text form, from `String()`: rank symbol plus suit symbol (`Q♣`). A Joker is `JK`; a zero Card is `??`. Always dependency-free, so `fmt.Println(card)` works without importing a renderer. | |

### Terms to resolve

Open questions where the right word isn't settled yet. Candidates for
`/grill-with-docs`:

- **Cut card** — the blank plastic marker in a blackjack shoe. Deck *state*
  (`CutAt(n)`) rather than a Card, to preserve deck homogeneity — but unverified
  against a real game.

## Deliberately avoided terms

| Avoid | Say instead | Why |
| --- | --- | --- |
| Pack | Deck | One word for the concept; "pack" invites a second type that behaves identically. |
| Shoe | Deck with copies | `NewDeck(WithCopies(6))`. A shoe draws and shuffles like any Deck, so a separate type would only duplicate it. |
| Hand, Pile, DiscardPile, Stock | Deck, or Way | A hand belongs to a *player*, and player state is out of scope. `Deck[T]` already does everything these need; each new name would be a near-copy of it. For what a Deal produces, say **way**. |
| Value (as a Card field) | Valuation | Value is a property of a card *in a game*, never of the card itself. Games supply a valuation function. |
| Points, Score, Beats, Trump | — | Game rules. gocard never decides which card wins. |
| Player, Seat, Round, Turn | — | Table state. Out of scope. |

## See also

- `docs/adr/` — architectural decisions
- `docs/agents/domain.md` — how skills consume this file
