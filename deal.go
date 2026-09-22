package gocard

import "fmt"

// Deal removes cards from the top of the deck and distributes them into
// ways — separate groupings of cardsPerWay cards each.
//
// Cards go one at a time to each way in turn, the order a dealer uses. That
// is observably different from taking cardsPerWay cards for each way in
// succession: dealing two ways of two from a fresh deck gives the first way
// the first and third cards, not the first two.
//
// Deal is atomic: if the deck holds fewer than ways*cardsPerWay cards it
// returns an error matching [ErrInsufficientCards] and removes nothing, so a
// short deal cannot leave some ways already filled. Dealing no ways, or no
// cards per way, is a no-op.
//
// gocard has no notion of who holds a way; a way is a slice of cards and
// nothing more.
func (d *Deck[T]) Deal(ways, cardsPerWay int) ([][]T, error) {
	if ways == 0 || cardsPerWay == 0 {
		return nil, nil
	}

	// Both arguments are checked here rather than left to DrawN, which sees only
	// their product and cannot tell a valid request from a wrapped one. Two
	// negatives multiply to a plausible positive, so Deal(-2, -2) would ask for
	// 4 cards; two large positives overflow to a plausible small positive, so
	// Deal(4, 2+(1<<62)) would ask for 8 and succeed before panicking on a
	// nonsensical allocation. Rejecting anything larger than the deck catches
	// both, since no deal can legitimately need more cards than the deck holds.
	if ways < 0 || cardsPerWay < 0 || ways > len(d.cards) || cardsPerWay > len(d.cards) {
		return nil, fmt.Errorf("dealing %d ways of %d: %w", ways, cardsPerWay, ErrInsufficientCards)
	}

	needed := ways * cardsPerWay

	drawn, err := d.DrawN(needed)
	if err != nil {
		return nil, fmt.Errorf("dealing %d ways of %d: %w", ways, cardsPerWay, err)
	}

	dealt := make([][]T, ways)
	for i := range dealt {
		dealt[i] = make([]T, 0, cardsPerWay)
	}

	for i, card := range drawn {
		way := i % ways
		dealt[way] = append(dealt[way], card)
	}

	return dealt, nil
}
