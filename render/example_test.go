package render_test

import (
	"fmt"

	"github.com/asphaltbuffet/gocard"
	"github.com/asphaltbuffet/gocard/render"
)

func ExampleCard() {
	c := gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}

	fmt.Println(render.Card(c, render.Plain{}))
	// Output:
	// ╭─────╮
	// │Q♣   │
	// │  ♣  │
	// │Q♣   │
	// ╰─────╯
}
