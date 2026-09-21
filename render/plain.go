package render

import (
	"strings"

	"github.com/asphaltbuffet/gocard"
)

// borderCells is the number of cells a border consumes across a dimension:
// one on each side.
const borderCells = 2

// border holds the drawing characters for one [gocard.BorderStyle].
type border struct {
	topLeft, topRight       string
	bottomLeft, bottomRight string
	horizontal, vertical    string
}

// Plain renders a card with box-drawing characters and no colour.
//
// Plain ignores [gocard.Layout.Accent] entirely: it is the renderer to use
// for output that must stay plain text, such as a log or a test fixture.
// Because a card's accent is a semantic role rather than a colour, ignoring it
// costs nothing but the colour itself.
//
// The zero value is ready to use.
type Plain struct{}

// Render draws l with box-drawing characters, clipping content that does not
// fit and padding rows that fall short, so every row is exactly l.Width
// cells wide and the result is exactly l.Height rows tall.
//
// A layout with no width or height renders as the empty string.
func (Plain) Render(l gocard.Layout) string {
	if l.Width <= 0 || l.Height <= 0 {
		return ""
	}

	b := borderFor(l.Border)

	if l.Border == gocard.BorderNone {
		return strings.Join(fit(strings.Split(l.Content, "\n"), l.Width, l.Height), "\n")
	}

	interiorWidth := l.Width - borderCells
	interiorHeight := l.Height - borderCells

	if interiorWidth <= 0 || interiorHeight <= 0 {
		return strings.Join(fit(strings.Split(l.Content, "\n"), l.Width, l.Height), "\n")
	}

	body := fit(strings.Split(l.Content, "\n"), interiorWidth, interiorHeight)

	rows := make([]string, 0, l.Height)
	rows = append(rows, b.topLeft+strings.Repeat(b.horizontal, interiorWidth)+b.topRight)

	for _, line := range body {
		rows = append(rows, b.vertical+line+b.vertical)
	}

	rows = append(rows, b.bottomLeft+strings.Repeat(b.horizontal, interiorWidth)+b.bottomRight)

	return strings.Join(rows, "\n")
}

// fit clips or pads lines so the result is exactly height rows of exactly
// width cells each.
func fit(lines []string, width, height int) []string {
	out := make([]string, height)

	for i := range out {
		line := ""
		if i < len(lines) {
			line = lines[i]
		}

		runes := []rune(line)
		if len(runes) > width {
			runes = runes[:width]
		}

		out[i] = string(runes) + strings.Repeat(" ", width-len(runes))
	}

	return out
}

// borderFor returns the drawing characters for style.
func borderFor(style gocard.BorderStyle) border {
	switch style {
	case gocard.BorderRounded:
		return border{"╭", "╮", "╰", "╯", "─", "│"}
	case gocard.BorderDouble:
		return border{"╔", "╗", "╚", "╝", "═", "║"}
	case gocard.BorderThick:
		return border{"┏", "┓", "┗", "┛", "━", "┃"}
	case gocard.BorderPlain:
		return border{"┌", "┐", "└", "┘", "─", "│"}
	case gocard.BorderNone:
		return border{" ", " ", " ", " ", " ", " "}
	}

	return border{"┌", "┐", "└", "┘", "─", "│"}
}
