# gocard

> A Go library for representing and manipulating playing cards and decks.

## Installation

```bash
go get github.com/asphaltbuffet/gocard
```

## Usage

```go
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/asphaltbuffet/gocard"
	"github.com/asphaltbuffet/gocard/render"
)

func main() {
	d := gocard.NewDeck()

	// Shuffle takes an explicit source, so a deal is reproducible when you
	// want it to be. Seed from rand.Uint64() for a different order each run.
	d.Shuffle(rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())))

	ways, err := d.Deal(4, 5)
	if err != nil {
		panic(err)
	}

	for _, card := range ways[0] {
		fmt.Println(render.Card(card, render.Plain{}))
	}
}
```

- **Cards** are comparable structs — `gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}`
  prints as `Q♣` with no renderer needed.
- **Decks** are generic over their card type, arrive ordered rather than
  shuffled, and draw and deal atomically.
- **Rendering** is layered: the core package has no external dependencies,
  `gocard/render` adds a plain renderer, and `gocard/render/lipgloss` adds
  colour. Arranging several cards is left to you — the boxes compose with
  `lipgloss.JoinHorizontal`.

A card carries no game value: what a card is worth depends on the game, so
supply a `gocard.Valuer` or use `gocard.PipValue` / `gocard.BlackjackValue`.

See [CONTEXT.md](CONTEXT.md) for the project's vocabulary and
[docs/adr/](docs/adr/) for why the API is shaped this way.

## Development

```bash
nix develop       # enter dev shell (Go toolchain + tooling)
mise run test     # test
mise run lint     # lint
mise run dev      # generate + lint + test
```

## License

MIT — see [LICENSE](LICENSE)
