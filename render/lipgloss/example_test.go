package lipgloss_test

import (
	"fmt"
	"strings"

	"github.com/asphaltbuffet/gocard"
	"github.com/asphaltbuffet/gocard/render"
	glrender "github.com/asphaltbuffet/gocard/render/lipgloss"
)

// ExampleNew renders a card and reports its row count. The raw box is not
// printed directly because its colour escape codes depend on the terminal's
// colour-profile detection, which is not deterministic under `go test`.
func ExampleNew() {
	c := gocard.Card{Rank: gocard.Queen, Suit: gocard.Hearts}

	out := render.Card(c, glrender.New())

	fmt.Println(len(strings.Split(out, "\n")))
	// Output: 5
}
