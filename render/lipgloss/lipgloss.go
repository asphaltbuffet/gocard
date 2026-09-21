// Package lipgloss renders gocard cards as styled terminal boxes.
//
// It implements [render.Renderer] on top of the lipgloss styling library,
// resolving a card's semantic [gocard.Accent] into an actual colour and its
// [gocard.BorderStyle] into a lipgloss border. Cards therefore never
// hard-code a palette: re-theme the renderer and every card follows.
//
// This is the only package in the module with an external dependency.
// Importing the core gocard package pulls nothing extra.
//
// Render returns a self-contained box, so a game author composes a display
// with lipgloss directly:
//
//	left := render.Card(one, r)
//	right := render.Card(two, r)
//	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
//
// gocard does not arrange multiple cards itself.
package lipgloss

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/asphaltbuffet/gocard"
)

// borderCells is the number of cells a border consumes across a dimension:
// one on each side.
const borderCells = 2

// Theme maps gocard's semantic accents to colours, as lipgloss colour
// strings such as "#ff0000" or an ANSI index like "1".
type Theme struct {
	// Red styles content a card marks [gocard.AccentRed], such as hearts.
	Red string

	// Black styles content a card marks [gocard.AccentBlack].
	Black string

	// Muted styles content a card marks [gocard.AccentMuted].
	Muted string

	// Border styles the card's frame.
	Border string
}

// DefaultTheme returns a theme that reads well on both light and dark
// terminals: adaptive-safe ANSI colours rather than fixed hex values.
func DefaultTheme() Theme {
	return Theme{
		Red:    "1",
		Black:  "7",
		Muted:  "8",
		Border: "8",
	}
}

// Option adjusts a [Renderer].
type Option func(*Renderer)

// WithTheme sets the colours the renderer resolves accents to.
func WithTheme(t Theme) Option {
	return func(r *Renderer) {
		r.theme = t
	}
}

// Renderer draws cards as styled lipgloss boxes.
//
// Use [New] to build one; the zero value has no theme and renders without
// colour.
type Renderer struct {
	theme Theme
}

// New returns a renderer using [DefaultTheme], adjusted by any options.
func New(opts ...Option) *Renderer {
	r := &Renderer{theme: DefaultTheme()}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// Render draws l as a styled box exactly l.Width cells wide and l.Height
// rows tall. A layout with no width or height renders as the empty string.
//
// Content is centred in the box, and trailing blank lines are dropped before
// centring, so content is positioned by how many non-blank lines it has rather
// than by where they sit in the string. Output is therefore not byte-comparable
// with the render package's Plain renderer, which left-aligns and keeps blank
// lines; both agree on the overall width and height.
func (r *Renderer) Render(l gocard.Layout) string {
	if l.Width <= 0 || l.Height <= 0 {
		return ""
	}

	style := lipgloss.NewStyle().
		Width(l.Width-borderCells).
		Height(l.Height-borderCells).
		Align(lipgloss.Center, lipgloss.Center).
		Foreground(lipgloss.Color(r.colorFor(l.Accent)))

	if l.Border != gocard.BorderNone {
		style = style.
			Border(borderFor(l.Border)).
			BorderForeground(lipgloss.Color(r.theme.Border))
	} else {
		style = style.Width(l.Width).Height(l.Height)
	}

	return style.Render(strings.TrimRight(l.Content, "\n"))
}

// colorFor resolves a semantic accent into a theme colour.
func (r *Renderer) colorFor(a gocard.Accent) string {
	switch a {
	case gocard.AccentRed:
		return r.theme.Red
	case gocard.AccentBlack:
		return r.theme.Black
	case gocard.AccentMuted:
		return r.theme.Muted
	case gocard.AccentNone:
		return ""
	}

	return ""
}

// borderFor resolves a gocard border style into a lipgloss border.
func borderFor(style gocard.BorderStyle) lipgloss.Border {
	switch style {
	case gocard.BorderRounded:
		return lipgloss.RoundedBorder()
	case gocard.BorderDouble:
		return lipgloss.DoubleBorder()
	case gocard.BorderThick:
		return lipgloss.ThickBorder()
	case gocard.BorderPlain:
		return lipgloss.NormalBorder()
	case gocard.BorderNone:
		return lipgloss.HiddenBorder()
	}

	return lipgloss.NormalBorder()
}
