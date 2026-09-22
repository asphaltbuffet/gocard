package gocard

import (
	"errors"
	"fmt"
)

// ErrInsufficientCards reports that an operation asked for more cards than
// the deck holds.
//
// Drawing and dealing are atomic: when this error is returned the deck is
// unchanged, so a short deal fails cleanly rather than leaving some cards
// already handed out. Match it with [errors.Is].
var ErrInsufficientCards = errors.New("insufficient cards")

// Draw removes the top card from the deck and returns it.
//
// The top of a deck is the card at index 0, so a freshly built deck draws the
// ace of clubs first. Drawing mutates the deck.
//
// If the deck is empty, Draw returns the zero card and an error matching
// [ErrInsufficientCards], and the deck is unchanged.
func (d *Deck[T]) Draw() (T, error) {
	cards, err := d.DrawN(1)
	if err != nil {
		var zero T

		return zero, err
	}

	return cards[0], nil
}

// DrawN removes the top n cards from the deck and returns them, top first.
//
// DrawN is atomic: if the deck holds fewer than n cards it returns an error
// matching [ErrInsufficientCards] and removes nothing. Drawing 0 cards is a
// no-op that returns no cards and no error; the returned slice is empty, which
// is to say safe to range over and append to, but not guaranteed non-nil.
func (d *Deck[T]) DrawN(n int) ([]T, error) {
	if n == 0 {
		return nil, nil
	}

	if n < 0 || n > len(d.cards) {
		return nil, fmt.Errorf("drawing %d of %d cards: %w", n, len(d.cards), ErrInsufficientCards)
	}

	drawn := make([]T, n)
	copy(drawn, d.cards[:n])
	d.cards = d.cards[n:]

	return drawn, nil
}
