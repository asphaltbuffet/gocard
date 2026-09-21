package gocard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asphaltbuffet/gocard"
)

func TestPipValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		card gocard.Card
		want int
	}{
		{"ace is low", gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs}, 1},
		{"seven", gocard.Card{Rank: gocard.Seven, Suit: gocard.Clubs}, 7},
		{"ten", gocard.Card{Rank: gocard.Ten, Suit: gocard.Clubs}, 10},
		{"jack continues the sequence", gocard.Card{Rank: gocard.Jack, Suit: gocard.Clubs}, 11},
		{"king is thirteen", gocard.Card{Rank: gocard.King, Suit: gocard.Clubs}, 13},
		{"joker has no pip value", gocard.Card{Rank: gocard.Joker}, 0},
		{"zero card", gocard.Card{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, gocard.PipValue(tt.card))
		})
	}
}

func TestBlackjackValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		card gocard.Card
		want int
	}{
		{"ace is high until the hand busts", gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs}, 11},
		{"seven", gocard.Card{Rank: gocard.Seven, Suit: gocard.Clubs}, 7},
		{"ten", gocard.Card{Rank: gocard.Ten, Suit: gocard.Clubs}, 10},
		{"jack counts ten", gocard.Card{Rank: gocard.Jack, Suit: gocard.Clubs}, 10},
		{"queen counts ten", gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}, 10},
		{"king counts ten", gocard.Card{Rank: gocard.King, Suit: gocard.Clubs}, 10},
		{"joker is not a blackjack card", gocard.Card{Rank: gocard.Joker}, 0},
		{"zero card", gocard.Card{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, gocard.BlackjackValue(tt.card))
		})
	}
}

func TestValuerAcceptsAFunction(t *testing.T) {
	t.Parallel()

	// A Valuer is just a function, so a caller can supply their own.
	var v gocard.Valuer[gocard.Card] = func(gocard.Card) int { return 42 }

	assert.Equal(t, 42, v(gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs}))
}
