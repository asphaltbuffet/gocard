package gocard_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asphaltbuffet/gocard"
)

func TestCardString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		card gocard.Card
		want string
	}{
		{"queen of clubs", gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}, "Q♣"},
		{"ace of spades", gocard.Card{Rank: gocard.Ace, Suit: gocard.Spades}, "A♠"},
		{"ten of hearts is two glyphs wide", gocard.Card{Rank: gocard.Ten, Suit: gocard.Hearts}, "10♥"},
		{"two of diamonds", gocard.Card{Rank: gocard.Two, Suit: gocard.Diamonds}, "2♦"},
		{"jack of hearts", gocard.Card{Rank: gocard.Jack, Suit: gocard.Hearts}, "J♥"},
		{"king of spades", gocard.Card{Rank: gocard.King, Suit: gocard.Spades}, "K♠"},
		{"joker has no suit", gocard.Card{Rank: gocard.Joker}, "JK"},
		{"zero card is unknown", gocard.Card{}, "??"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.card.String())
		})
	}
}

func TestCardStringValueAndPointerAgree(t *testing.T) {
	t.Parallel()

	c := gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}

	assert.Equal(t, fmt.Sprint(c), fmt.Sprint(&c),
		"String must be on the value receiver so both forms print alike")
}

func TestCardComparable(t *testing.T) {
	t.Parallel()

	a := gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}
	b := gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}
	c := gocard.Card{Rank: gocard.Queen, Suit: gocard.Hearts}

	assert.Equal(t, a, b, "identical cards must compare equal")
	assert.NotEqual(t, a, c, "different suits must not compare equal")

	seen := map[gocard.Card]bool{a: true}
	assert.True(t, seen[b], "Card must be usable as a map key")
}

func TestCardIsJoker(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		card gocard.Card
		want bool
	}{
		{"joker", gocard.Card{Rank: gocard.Joker}, true},
		{"queen of clubs", gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}, false},
		{"zero card", gocard.Card{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.card.IsJoker())
		})
	}
}

func TestCardValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		card gocard.Card
		want bool
	}{
		{"queen of clubs", gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}, true},
		{"joker needs no suit", gocard.Card{Rank: gocard.Joker}, true},
		{"zero card", gocard.Card{}, false},
		{"rank without suit", gocard.Card{Rank: gocard.Queen}, false},
		{"suit without rank", gocard.Card{Suit: gocard.Clubs}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.card.Valid())
		})
	}
}
