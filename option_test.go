package gocard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/gocard"
)

func TestWithJokers(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck(gocard.WithJokers(2))

	require.Equal(t, standardDeckSize+2, d.Len(), "jokers are added to the 52")

	jokers := 0

	for _, c := range d.Cards() {
		if c.IsJoker() {
			jokers++

			assert.Equal(t, gocard.NoSuit, c.Suit, "a joker has no suit")
		}
	}

	assert.Equal(t, 2, jokers)
}

func TestWithJokersNegativeIsIgnored(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck(gocard.WithJokers(-3))

	assert.Equal(t, standardDeckSize, d.Len(), "a negative count clamps to zero")
}

func TestWithRanksBuildsAEuchreDeck(t *testing.T) {
	t.Parallel()

	// Euchre uses nine through ace in four suits: 24 cards.
	d := gocard.NewDeck(gocard.WithRanks(
		gocard.Nine, gocard.Ten, gocard.Jack, gocard.Queen, gocard.King, gocard.Ace,
	))

	require.Equal(t, 24, d.Len())

	for _, c := range d.Cards() {
		assert.NotEqual(t, gocard.Two, c.Rank, "low ranks are excluded")
	}
}

func TestWithRanksEmptyIsIgnored(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck(gocard.WithRanks())

	assert.Equal(t, standardDeckSize, d.Len(),
		"an empty rank list falls back to the standard ranks")
}

func TestWithSuitsEmptyIsIgnored(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck(gocard.WithSuits())

	assert.Equal(t, standardDeckSize, d.Len(),
		"an empty suit list falls back to the standard suits")
}

func TestWithSuits(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck(gocard.WithSuits(gocard.Hearts, gocard.Spades))

	require.Equal(t, 26, d.Len())

	for _, c := range d.Cards() {
		assert.Contains(t, []gocard.Suit{gocard.Hearts, gocard.Spades}, c.Suit)
	}
}

func TestWithCopiesBuildsAShoe(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck(gocard.WithCopies(6))

	assert.Equal(t, standardDeckSize*6, d.Len(),
		"six copies is what a casino calls a shoe; it is still just a Deck")
}

func TestWithCopiesBelowOneIsClamped(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck(gocard.WithCopies(0))

	assert.Equal(t, standardDeckSize, d.Len(), "a count below one clamps to one deck")
}

func TestWithCopiesRepeatsJokers(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck(gocard.WithCopies(2), gocard.WithJokers(1))

	assert.Equal(t, (standardDeckSize+1)*2, d.Len(),
		"each copy brings its own jokers")
}

func TestOptionsCombine(t *testing.T) {
	t.Parallel()

	// A two-deck pinochle-style pack: nine through ace, two copies.
	d := gocard.NewDeck(
		gocard.WithRanks(gocard.Nine, gocard.Ten, gocard.Jack, gocard.Queen, gocard.King, gocard.Ace),
		gocard.WithCopies(2),
	)

	assert.Equal(t, 48, d.Len())
}
