package gocard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/gocard"
)

const standardDeckSize = 52

func TestNewDeckHasFiftyTwoCards(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	assert.Equal(t, standardDeckSize, d.Len(), "a standard deck has 52 cards and no jokers")
}

func TestNewDeckIsOrderedNotShuffled(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()
	cards := d.Cards()

	require.Len(t, cards, standardDeckSize)

	assert.Equal(t, gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs}, cards[0],
		"a new deck starts with the ace of clubs")
	assert.Equal(t, gocard.Card{Rank: gocard.King, Suit: gocard.Clubs}, cards[12],
		"clubs run ace through king before diamonds begin")
	assert.Equal(t, gocard.Card{Rank: gocard.Ace, Suit: gocard.Diamonds}, cards[13],
		"diamonds follow clubs")
	assert.Equal(t, gocard.Card{Rank: gocard.King, Suit: gocard.Spades}, cards[51],
		"a new deck ends with the king of spades")
}

func TestNewDeckHasNoDuplicates(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	seen := make(map[gocard.Card]bool, standardDeckSize)
	for _, c := range d.Cards() {
		require.False(t, seen[c], "card %s appears twice", c)
		seen[c] = true
	}

	assert.Len(t, seen, standardDeckSize)
}

func TestNewDeckIsReproducible(t *testing.T) {
	t.Parallel()

	first := gocard.NewDeck().Cards()
	second := gocard.NewDeck().Cards()

	assert.Equal(t, first, second, "construction must be deterministic so tests can rely on it")
}

func TestDeckCardsReturnsACopy(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	cards := d.Cards()
	cards[0] = gocard.Card{Rank: gocard.Joker}

	assert.Equal(t, gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs}, d.Cards()[0],
		"mutating the returned slice must not reach into the deck")
}

func TestDeckCloneStartsEqual(t *testing.T) {
	t.Parallel()

	original := gocard.NewDeck()
	clone := original.Clone()

	assert.Equal(t, original.Cards(), clone.Cards(), "a clone starts equal")
	assert.Equal(t, original.Len(), clone.Len())
}

func TestDeckCloneDoesNotShareBackingArray(t *testing.T) {
	t.Parallel()

	original := gocard.NewDeck()
	clone := original.Clone()

	// Reach in through the copy Cards returns to prove the clone's storage is
	// its own. Independence under Draw is covered once Draw exists.
	cards := clone.Cards()
	cards[0] = gocard.Card{Rank: gocard.Joker}

	assert.Equal(t, gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs}, original.Cards()[0])
	assert.Equal(t, gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs}, clone.Cards()[0])
}

func TestNewDeckOfHoldsArbitraryCards(t *testing.T) {
	t.Parallel()

	type tarot struct{ Name string }

	d := gocard.NewDeckOf(tarot{"The Fool"}, tarot{"The Magician"})

	assert.Equal(t, 2, d.Len(), "Deck is generic over the card type")
}

func TestDeckZeroValueIsEmpty(t *testing.T) {
	t.Parallel()

	var d gocard.Deck[gocard.Card]

	assert.Equal(t, 0, d.Len(), "the zero Deck is an empty deck, not a broken one")
	assert.Empty(t, d.Cards())
}
