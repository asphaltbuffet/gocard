// Package render draws gocard cards as terminal output.
//
// A card describes itself as a [gocard.Layout]; a Renderer turns that
// description into text. The two never meet directly, so a renderer written
// today can draw a card family invented tomorrow, and the core gocard package
// needs no rendering library.
//
// [Plain] renders with box-drawing characters and no colour, and needs no
// dependencies. For colour, see the gocard/render/lipgloss package.
package render

import "github.com/asphaltbuffet/gocard"

// Renderer turns a card's layout into terminal output.
//
// This interface lives here rather than in gocard so that the core package
// does not depend on rendering, and so third parties can supply a renderer
// without gocard exporting anything for them.
type Renderer interface {
	// Render draws l and returns the result, with "\n" between rows.
	Render(l gocard.Layout) string
}

// Layouter is anything that can describe how it presents itself — every card
// family, including card types defined outside gocard.
type Layouter interface {
	// Layout returns the card's presentation description.
	Layout() gocard.Layout
}

// Card draws c using r.
//
// It accepts any type with a Layout method, so it works for
// [gocard.Card] and for card families this package has never seen:
//
//	fmt.Println(render.Card(card, render.Plain{}))
func Card(c Layouter, r Renderer) string {
	return r.Render(c.Layout())
}
