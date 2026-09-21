package gocard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asphaltbuffet/gocard"
)

func TestRankZeroValueIsNoRank(t *testing.T) {
	t.Parallel()

	var r gocard.Rank

	assert.Equal(t, gocard.NoRank, r, "zero Rank must mean absent, not Ace")
}

func TestRankOrdinal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		rank gocard.Rank
		want int
	}{
		{"no rank has no ordinal", gocard.NoRank, 0},
		{"ace is low by convention", gocard.Ace, 1},
		{"ten", gocard.Ten, 10},
		{"jack", gocard.Jack, 11},
		{"queen", gocard.Queen, 12},
		{"king is high", gocard.King, 13},
		{"joker has no ordinal", gocard.Joker, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.rank.Ordinal())
		})
	}
}

func TestRankString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		rank gocard.Rank
		want string
	}{
		{"ace", gocard.Ace, "Ace"},
		{"ten", gocard.Ten, "Ten"},
		{"king", gocard.King, "King"},
		{"joker", gocard.Joker, "Joker"},
		{"no rank", gocard.NoRank, "NoRank"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.rank.String())
		})
	}
}

func TestRankValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		rank gocard.Rank
		want bool
	}{
		{"no rank is invalid", gocard.NoRank, false},
		{"ace is valid", gocard.Ace, true},
		{"king is valid", gocard.King, true},
		{"joker is valid", gocard.Joker, true},
		{"out of range is invalid", gocard.Rank(200), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.rank.Valid())
		})
	}
}
