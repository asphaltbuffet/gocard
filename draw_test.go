package gocard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/gocard"
)

func TestDrawTakesFromTheTop(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	c, err := d.Draw()
	require.NoError(t, err)

	assert.Equal(t, gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs}, c,
		"the top of the deck is index 0")
	assert.Equal(t, standardDeckSize-1, d.Len(), "Draw mutates the deck")
}

func TestDrawEmptyDeckReportsInsufficientCards(t *testing.T) {
	t.Parallel()

	var d gocard.Deck[gocard.Card]

	c, err := d.Draw()

	require.ErrorIs(t, err, gocard.ErrInsufficientCards)
	assert.Equal(t, gocard.Card{}, c, "the zero card comes back with the error")
	assert.Equal(t, 0, d.Len())
}

func TestDrawNTakesInOrder(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	cards, err := d.DrawN(3)
	require.NoError(t, err)

	assert.Equal(t, []gocard.Card{
		{Rank: gocard.Ace, Suit: gocard.Clubs},
		{Rank: gocard.Two, Suit: gocard.Clubs},
		{Rank: gocard.Three, Suit: gocard.Clubs},
	}, cards)
	assert.Equal(t, standardDeckSize-3, d.Len())
}

func TestDrawNIsAtomic(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeckOf(
		gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs},
		gocard.Card{Rank: gocard.Two, Suit: gocard.Clubs},
		gocard.Card{Rank: gocard.Three, Suit: gocard.Clubs},
	)

	cards, err := d.DrawN(5)

	require.ErrorIs(t, err, gocard.ErrInsufficientCards)
	assert.Nil(t, cards, "an atomic failure returns no cards")
	assert.Equal(t, 3, d.Len(), "a failed draw must remove nothing")
}

func TestDrawNZeroIsANoOp(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	cards, err := d.DrawN(0)

	require.NoError(t, err)
	assert.Empty(t, cards)
	assert.Equal(t, standardDeckSize, d.Len())
}

func TestDrawNNegativeReportsInsufficientCards(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	cards, err := d.DrawN(-1)

	require.ErrorIs(t, err, gocard.ErrInsufficientCards)
	assert.Nil(t, cards)
	assert.Equal(t, standardDeckSize, d.Len(), "the deck is untouched")
}

func TestDrawNExactlyEmptiesTheDeck(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	cards, err := d.DrawN(standardDeckSize)

	require.NoError(t, err)
	assert.Len(t, cards, standardDeckSize)
	assert.Equal(t, 0, d.Len())
}

func TestDrawErrorIsMatchableWithErrorsIs(t *testing.T) {
	t.Parallel()

	var d gocard.Deck[gocard.Card]

	_, err := d.Draw()

	require.Error(t, err)
	assert.ErrorIs(t, err, gocard.ErrInsufficientCards,
		"callers branch on the sentinel, so it must survive wrapping")
}

func TestDrawFromCloneLeavesOriginalIntact(t *testing.T) {
	t.Parallel()

	original := gocard.NewDeck()
	clone := original.Clone()

	_, err := clone.Draw()
	require.NoError(t, err)

	assert.Equal(t, standardDeckSize, original.Len(),
		"drawing from a clone must not affect the original")
	assert.Equal(t, standardDeckSize-1, clone.Len())
}
