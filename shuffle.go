package gocard

import "math/rand/v2"

// Shuffle reorders the deck in place using r.
//
// The random source is required rather than optional, so that a shuffle is
// reproducible by default: passing a seeded generator is what lets a game's
// tests assert on a known deal.
//
//	d.Shuffle(rand.New(rand.NewPCG(seed, seed)))
//
// For a different order on every run, seed from the default source:
//
//	d.Shuffle(rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())))
//
// A nil r leaves the deck unchanged.
func (d *Deck[T]) Shuffle(r *rand.Rand) {
	if r == nil {
		return
	}

	r.Shuffle(len(d.cards), d.swap)
}

// swap exchanges the cards at i and j. It is the swapper Shuffle hands to the
// standard library.
func (d *Deck[T]) swap(i, j int) {
	d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
}
