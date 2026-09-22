package gocard_test

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/gocard"
)

// seed is an arbitrary fixed value; any constant makes the test reproducible.
const seed = 1

func newSeeded() *rand.Rand {
	return rand.New(rand.NewPCG(seed, seed))
}

func TestShuffleIsDeterministic(t *testing.T) {
	t.Parallel()

	first := gocard.NewDeck()
	first.Shuffle(newSeeded())

	second := gocard.NewDeck()
	second.Shuffle(newSeeded())

	assert.Equal(t, first.Cards(), second.Cards(),
		"the same seed must produce the same order, or games cannot be tested")
}

func TestShuffleChangesOrder(t *testing.T) {
	t.Parallel()

	ordered := gocard.NewDeck().Cards()

	d := gocard.NewDeck()
	d.Shuffle(newSeeded())

	assert.NotEqual(t, ordered, d.Cards(), "shuffling must reorder the deck")
}

func TestShuffleDiffersBetweenSeeds(t *testing.T) {
	t.Parallel()

	first := gocard.NewDeck()
	first.Shuffle(rand.New(rand.NewPCG(1, 1)))

	second := gocard.NewDeck()
	second.Shuffle(rand.New(rand.NewPCG(2, 2)))

	assert.NotEqual(t, first.Cards(), second.Cards(),
		"different seeds must produce different orders")
}

func TestShufflePreservesEveryCard(t *testing.T) {
	t.Parallel()

	d := gocard.NewDeck()
	d.Shuffle(newSeeded())

	require.Equal(t, standardDeckSize, d.Len(), "shuffling must not lose cards")

	seen := make(map[gocard.Card]bool, standardDeckSize)
	for _, c := range d.Cards() {
		require.False(t, seen[c], "card %s appears twice after shuffling", c)
		seen[c] = true
	}

	assert.Len(t, seen, standardDeckSize)
}

func TestShuffleEmptyDeckIsSafe(t *testing.T) {
	t.Parallel()

	var d gocard.Deck[gocard.Card]

	d.Shuffle(newSeeded())

	assert.Equal(t, 0, d.Len(), "shuffling an empty deck must not panic")
}

func TestShuffleNilSourceIsSafe(t *testing.T) {
	t.Parallel()

	ordered := gocard.NewDeck().Cards()

	d := gocard.NewDeck()
	d.Shuffle(nil)

	assert.Equal(t, ordered, d.Cards(),
		"a nil source leaves the deck untouched rather than panicking")
}
