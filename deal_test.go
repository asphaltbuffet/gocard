package gocard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/gocard"
)

func TestDealIsRoundRobin(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	ways, err := d.Deal(2, 2)
	require.NoError(t, err)
	require.Len(t, ways, 2)

	// Cards go A♣ 2♣ 3♣ 4♣ one at a time to each way in turn, so the first
	// way gets the first and third cards, not the first two.
	assert.Equal(t, []gocard.Card{
		{Rank: gocard.Ace, Suit: gocard.Clubs},
		{Rank: gocard.Three, Suit: gocard.Clubs},
	}, ways[0], "a dealer alternates; this is not DrawN(2)")

	assert.Equal(t, []gocard.Card{
		{Rank: gocard.Two, Suit: gocard.Clubs},
		{Rank: gocard.Four, Suit: gocard.Clubs},
	}, ways[1])
}

func TestDealRemovesEveryCardDealt(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	_, err := d.Deal(4, 5)
	require.NoError(t, err)

	assert.Equal(t, standardDeckSize-20, d.Len())
}

func TestDealShapesTheResult(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	ways, err := d.Deal(3, 7)
	require.NoError(t, err)

	require.Len(t, ways, 3, "the outer slice is one entry per way")

	for i, way := range ways {
		assert.Len(t, way, 7, "way %d holds cardsPerWay cards", i)
	}
}

func TestDealIsAtomic(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeckOf(
		gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs},
		gocard.Card{Rank: gocard.Two, Suit: gocard.Clubs},
		gocard.Card{Rank: gocard.Three, Suit: gocard.Clubs},
	)

	ways, err := d.Deal(2, 2)

	require.ErrorIs(t, err, gocard.ErrInsufficientCards)
	assert.Nil(t, ways, "a short deal returns nothing")
	assert.Equal(t, 3, d.Len(), "a short deal must leave the deck untouched")
}

func TestDealExactlyEmptiesTheDeck(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	ways, err := d.Deal(4, 13)

	require.NoError(t, err)
	assert.Len(t, ways, 4)
	assert.Equal(t, 0, d.Len())
}

func TestDealZeroWaysOrCards(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		ways        int
		cardsPerWay int
	}{
		{"no ways", 0, 5},
		{"no cards per way", 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := gocard.NewDeck()

			ways, err := d.Deal(tt.ways, tt.cardsPerWay)

			require.NoError(t, err)
			assert.Empty(t, ways)
			assert.Equal(t, standardDeckSize, d.Len(), "nothing is dealt")
		})
	}
}

func TestDealNegativeReportsInsufficientCards(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()

	ways, err := d.Deal(-1, 5)

	require.ErrorIs(t, err, gocard.ErrInsufficientCards)
	assert.Nil(t, ways)
	assert.Equal(t, standardDeckSize, d.Len())
}
