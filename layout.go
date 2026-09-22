package gocard

import "strings"

//go:generate stringer -type=BorderStyle -trimprefix=
//go:generate stringer -type=Accent -trimprefix=

// Default dimensions of a playing card's box, in terminal cells. The width
// accommodates the widest glyph ("10♥") with a border and padding on each
// side.
const (
	defaultCardWidth  = 7
	defaultCardHeight = 5
)

// BorderStyle names the frame a card asks to be drawn with.
//
// It is a style, not a set of characters: each renderer decides what a
// rounded border looks like in its output mode, and a plain-text renderer may
// ignore the distinction entirely.
type BorderStyle uint8

// Border styles a [Layout] can request. BorderNone is the zero value.
const (
	BorderNone BorderStyle = iota
	BorderPlain
	BorderRounded
	BorderDouble
	BorderThick
)

// Accent is a semantic colour role on a [Layout], not a colour.
//
// AccentRed means "this is a red suit", leaving the renderer's theme to
// decide what red looks like — or to ignore colour altogether. Cards
// therefore never hard-code a palette, and a colour scheme can change without
// touching any card.
type Accent uint8

// Accent roles a [Layout] can carry. AccentNone is the zero value.
const (
	AccentNone Accent = iota
	AccentRed
	AccentBlack
	AccentMuted
)

// Layout is a card's own description of how it presents itself: a single
// bordered box of a given size holding unstyled content.
//
// A card builds its Layout from this package's types alone, so the core of
// gocard needs no rendering library. A renderer turns a Layout into terminal
// output; see the gocard/render package. Because a card controls its own size
// and border, a new card family displays correctly through renderers written
// before it existed. See
// docs/adr/0003-cards-describe-layout-renderers-draw-it.md.
//
// The zero value describes an empty, borderless, unaccented box.
type Layout struct {
	// Width and Height are the card's size in terminal cells, borders
	// included.
	Width, Height int

	// Border is the frame the card asks for.
	Border BorderStyle

	// Content is the card's text, unstyled, with "\n" separating lines. It
	// must never contain escape sequences: applying style is the renderer's
	// job.
	Content string

	// Accent is the semantic colour role for the card's content.
	Accent Accent
}

// Layout returns how the card presents itself: a rounded box with the card's
// glyph centred, accented by suit colour.
//
// The content is unstyled text. Pass the result to a renderer — see
// gocard/render — or read it to build a display of your own.
func (c Card) Layout() Layout {
	return Layout{
		Width:   defaultCardWidth,
		Height:  defaultCardHeight,
		Border:  BorderRounded,
		Content: c.layoutContent(),
		Accent:  c.accent(),
	}
}

// layoutContent returns the card's text arranged for its box: the glyph at
// the top left, repeated bottom right, with the suit pip centred.
func (c Card) layoutContent() string {
	glyph := c.String()

	// Interior width and height exclude the border on each side; centring
	// splits the leftover space in two.
	const (
		borderCells = 2
		centreSplit = 2
	)

	interior := defaultCardWidth - borderCells
	rows := defaultCardHeight - borderCells

	lines := make([]string, 0, rows)
	lines = append(lines, glyph)

	// The centre shows a suit pip only when the glyph names a suited card.
	// String drops the suit for a joker or an absent rank, so presenting one
	// here would make the card print as one thing and render as another.
	middle := ""
	if !c.IsJoker() && c.Rank.Valid() {
		middle = c.Suit.Symbol()
	}

	if middle == "" {
		middle = glyph
	}

	pad := max((interior-len([]rune(middle)))/centreSplit, 0)

	lines = append(lines, strings.Repeat(" ", pad)+middle)
	lines = append(lines, glyph)

	return strings.Join(lines, "\n")
}

// accent returns the semantic colour role for the card's suit.
//
// A joker and a card with no rank carry no suit colour, matching [Card.String],
// which drops the suit for both.
func (c Card) accent() Accent {
	if c.IsJoker() || !c.Rank.Valid() {
		return AccentNone
	}

	switch c.Suit {
	case Hearts, Diamonds:
		return AccentRed
	case Clubs, Spades:
		return AccentBlack
	case NoSuit:
		return AccentNone
	}

	return AccentNone
}
