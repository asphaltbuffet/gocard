package gocard_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/gocard"
)

func TestCardLayoutHasUsableDimensions(t *testing.T) {
	t.Parallel()

	l := gocard.Card{Rank: gocard.Ten, Suit: gocard.Hearts}.Layout()

	assert.Positive(t, l.Width, "a card needs a width to be drawn")
	assert.Positive(t, l.Height, "a card needs a height to be drawn")
	assert.GreaterOrEqual(t, l.Width, len("10♥"),
		"the widest glyph must fit inside the card")
}

func TestCardLayoutContentCarriesTheGlyph(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		card gocard.Card
		want string
	}{
		{"queen of clubs", gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}, "Q♣"},
		{"ten of hearts", gocard.Card{Rank: gocard.Ten, Suit: gocard.Hearts}, "10♥"},
		{"joker", gocard.Card{Rank: gocard.Joker}, "JK"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := tt.card.Layout()

			assert.Contains(t, l.Content, tt.want)
		})
	}
}

func TestCardLayoutContentIsUnstyled(t *testing.T) {
	t.Parallel()

	l := gocard.Card{Rank: gocard.Queen, Suit: gocard.Hearts}.Layout()

	assert.NotContains(t, l.Content, "\x1b",
		"content must carry no escape sequences; the renderer applies style")
}

func TestCardLayoutAccentIsSemantic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		card gocard.Card
		want gocard.Accent
	}{
		{"hearts are red", gocard.Card{Rank: gocard.Ace, Suit: gocard.Hearts}, gocard.AccentRed},
		{"diamonds are red", gocard.Card{Rank: gocard.Ace, Suit: gocard.Diamonds}, gocard.AccentRed},
		{"clubs are black", gocard.Card{Rank: gocard.Ace, Suit: gocard.Clubs}, gocard.AccentBlack},
		{"spades are black", gocard.Card{Rank: gocard.Ace, Suit: gocard.Spades}, gocard.AccentBlack},
		{"a joker has no suit colour", gocard.Card{Rank: gocard.Joker}, gocard.AccentNone},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.card.Layout().Accent)
		})
	}
}

func TestLayoutZeroValue(t *testing.T) {
	t.Parallel()

	var l gocard.Layout

	assert.Equal(t, gocard.BorderNone, l.Border, "the zero border is none")
	assert.Equal(t, gocard.AccentNone, l.Accent, "the zero accent is none")
	assert.Equal(t, 0, l.Width)
	assert.Empty(t, l.Content)
}

func TestCardLayoutContentFitsTheBox(t *testing.T) {
	t.Parallel()

	// A border consumes one cell on each side of both dimensions, so content
	// must fit the interior, not the outer size. Checking against the interior
	// is what catches a change to the default width or height that the content
	// builder was not updated for.
	const borderCells = 2

	tests := []struct {
		name string
		card gocard.Card
	}{
		{"widest glyph", gocard.Card{Rank: gocard.Ten, Suit: gocard.Diamonds}},
		{"single-rune rank", gocard.Card{Rank: gocard.Queen, Suit: gocard.Hearts}},
		{"no suit falls back to the glyph", gocard.Card{Rank: gocard.Joker}},
		{"zero card", gocard.Card{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := tt.card.Layout()
			lines := strings.Split(l.Content, "\n")

			require.LessOrEqual(t, len(lines), l.Height-borderCells,
				"content must fit the box interior, not just the outer height")

			for i, line := range lines {
				assert.LessOrEqual(t, len([]rune(line)), l.Width-borderCells,
					"row %d must fit the interior width, measured in runes", i)
			}
		})
	}
}

func TestBorderStyleString(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "BorderRounded", gocard.BorderRounded.String())
	assert.Equal(t, "BorderNone", gocard.BorderNone.String())
}

func TestAccentString(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "AccentRed", gocard.AccentRed.String())
	assert.Equal(t, "AccentNone", gocard.AccentNone.String())
}
