package gocard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asphaltbuffet/gocard"
)

func TestSuitZeroValueIsNoSuit(t *testing.T) {
	t.Parallel()

	var s gocard.Suit

	assert.Equal(t, gocard.NoSuit, s, "zero Suit must mean absent, not Clubs")
}

func TestSuitSymbol(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		suit gocard.Suit
		want string
	}{
		{"no suit has no symbol", gocard.NoSuit, ""},
		{"clubs", gocard.Clubs, "♣"},
		{"diamonds", gocard.Diamonds, "♦"},
		{"hearts", gocard.Hearts, "♥"},
		{"spades", gocard.Spades, "♠"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.suit.Symbol())
		})
	}
}

func TestSuitValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		suit gocard.Suit
		want bool
	}{
		{"no suit is invalid", gocard.NoSuit, false},
		{"clubs is valid", gocard.Clubs, true},
		{"spades is valid", gocard.Spades, true},
		{"out of range is invalid", gocard.Suit(200), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.suit.Valid())
		})
	}
}
