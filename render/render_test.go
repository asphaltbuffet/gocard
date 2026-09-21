package render_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/gocard"
	"github.com/asphaltbuffet/gocard/render"
)

func TestPlainRendersABox(t *testing.T) {
	t.Parallel()

	out := render.Card(gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}, render.Plain{})

	lines := strings.Split(out, "\n")

	require.Len(t, lines, 5, "a standard card is five rows tall")

	for i, line := range lines {
		assert.Len(t, []rune(line), 7, "row %d must be the card's width", i)
	}
}

func TestPlainIncludesTheGlyph(t *testing.T) {
	t.Parallel()

	out := render.Card(gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}, render.Plain{})

	assert.Contains(t, out, "Q♣")
}

func TestPlainEmitsNoEscapeSequences(t *testing.T) {
	t.Parallel()

	out := render.Card(gocard.Card{Rank: gocard.Ace, Suit: gocard.Hearts}, render.Plain{})

	assert.NotContains(t, out, "\x1b",
		"the plain renderer ignores Accent entirely, which is the point of a semantic accent")
}

func TestPlainRendersEveryBorderStyle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		style gocard.BorderStyle
	}{
		{"none", gocard.BorderNone},
		{"plain", gocard.BorderPlain},
		{"rounded", gocard.BorderRounded},
		{"double", gocard.BorderDouble},
		{"thick", gocard.BorderThick},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := gocard.Layout{Width: 7, Height: 3, Border: tt.style, Content: "Q♣"}

			out := render.Plain{}.Render(l)

			lines := strings.Split(out, "\n")
			assert.Len(t, lines, 3, "every border style must respect the requested height")
		})
	}
}

func TestPlainHandlesTheZeroLayout(t *testing.T) {
	t.Parallel()

	out := render.Plain{}.Render(gocard.Layout{})

	assert.Empty(t, out, "a zero-sized layout renders nothing rather than panicking")
}

func TestPlainClipsOverlongContent(t *testing.T) {
	t.Parallel()

	l := gocard.Layout{
		Width:   7,
		Height:  3,
		Border:  gocard.BorderPlain,
		Content: "this line is far too long to fit\nand\nso\nare\nthese",
	}

	out := render.Plain{}.Render(l)

	lines := strings.Split(out, "\n")

	require.Len(t, lines, 3, "content taller than the box is clipped, not overflowed")

	for i, line := range lines {
		assert.Len(t, []rune(line), 7, "row %d must be clipped to the width", i)
	}
}

// customCard is a card family defined outside gocard, proving a renderer
// written today draws a card type it has never heard of.
type customCard struct{}

func (customCard) Layout() gocard.Layout {
	return gocard.Layout{Width: 6, Height: 3, Border: gocard.BorderDouble, Content: "WILD"}
}

func TestCardAcceptsAnyLayouter(t *testing.T) {
	t.Parallel()

	out := render.Card(customCard{}, render.Plain{})

	assert.Contains(t, out, "WILD",
		"render.Card takes a one-method interface, so third-party cards work")
}
