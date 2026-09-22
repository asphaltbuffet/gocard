package lipgloss_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/gocard"
	"github.com/asphaltbuffet/gocard/render"
	glrender "github.com/asphaltbuffet/gocard/render/lipgloss"
)

func TestRendererSatisfiesTheContract(t *testing.T) {
	t.Parallel()

	var r render.Renderer = glrender.New()

	assert.NotNil(t, r, "the lipgloss renderer must satisfy render.Renderer")
}

func TestRendersACardOfTheRequestedSize(t *testing.T) {
	t.Parallel()

	out := render.Card(gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}, glrender.New())

	lines := strings.Split(out, "\n")

	require.Len(t, lines, 5, "a standard card is five rows tall")
	assert.Contains(t, out, "Q♣")
}

func TestRendersEveryBorderStyle(t *testing.T) {
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

			l := gocard.Layout{Width: 7, Height: 5, Border: tt.style, Content: "Q♣"}

			out := glrender.New().Render(l)

			assert.NotEmpty(t, out, "every border style must render")
			assert.Contains(t, out, "Q♣")
		})
	}
}

func TestRendersEveryAccent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		accent gocard.Accent
	}{
		{"none", gocard.AccentNone},
		{"red", gocard.AccentRed},
		{"black", gocard.AccentBlack},
		{"muted", gocard.AccentMuted},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := gocard.Layout{Width: 7, Height: 5, Border: gocard.BorderRounded, Content: "Q♣", Accent: tt.accent}

			out := glrender.New().Render(l)

			assert.Contains(t, out, "Q♣", "every accent must still render the content")
		})
	}
}

func TestRendersExactlyTheDeclaredSize(t *testing.T) {
	t.Parallel()

	// A lipgloss style's Width and Height are minimums: it wraps and grows
	// rather than truncating. A card family with taller or wordier content than
	// a standard playing card would otherwise get a box the wrong size, which
	// breaks any caller joining boxes side by side.
	const (
		width  = 7
		height = 5
	)

	tests := []struct {
		name    string
		border  gocard.BorderStyle
		content string
	}{
		{"too many lines, bordered", gocard.BorderRounded, "a\nb\nc\nd\ne\nf\ng"},
		{"too many lines, borderless", gocard.BorderNone, "a\nb\nc\nd\ne\nf\ng"},
		{"unbreakable long word, bordered", gocard.BorderRounded, "unbreakableverylongword"},
		{"unbreakable long word, borderless", gocard.BorderNone, "unbreakableverylongword"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := gocard.Layout{Width: width, Height: height, Border: tt.border, Content: tt.content}

			lines := strings.Split(glrender.New().Render(l), "\n")

			require.Len(t, lines, height, "the box must be exactly as tall as the layout declares")

			for i, line := range lines {
				assert.Len(t, []rune(line), width,
					"row %d must be exactly the declared width in runes", i)
			}
		})
	}
}

func TestZeroLayoutRendersNothing(t *testing.T) {
	t.Parallel()

	assert.Empty(t, glrender.New().Render(gocard.Layout{}),
		"a zero-sized layout renders nothing rather than panicking")
}

func TestWithThemeOverridesColours(t *testing.T) {
	t.Parallel()

	theme := glrender.DefaultTheme()
	theme.Red = "#ff0000"

	r := glrender.New(glrender.WithTheme(theme))

	l := gocard.Layout{Width: 7, Height: 5, Border: gocard.BorderRounded, Content: "A♥", Accent: gocard.AccentRed}

	assert.Contains(t, r.Render(l), "A♥", "a custom theme must still render content")
}

func TestDefaultThemeIsPopulated(t *testing.T) {
	t.Parallel()

	theme := glrender.DefaultTheme()

	assert.NotEmpty(t, theme.Red)
	assert.NotEmpty(t, theme.Black)
	assert.NotEmpty(t, theme.Muted)
	assert.NotEmpty(t, theme.Border)
}

// customCard is a card family defined outside gocard.
type customCard struct{}

func (customCard) Layout() gocard.Layout {
	return gocard.Layout{
		Width:   9,
		Height:  5,
		Border:  gocard.BorderDouble,
		Content: "WILD",
		Accent:  gocard.AccentMuted,
	}
}

func TestRendersAThirdPartyCardFamily(t *testing.T) {
	t.Parallel()

	out := render.Card(customCard{}, glrender.New())

	assert.Contains(t, out, "WILD",
		"a renderer written today must draw a card family it has never heard of")
}
