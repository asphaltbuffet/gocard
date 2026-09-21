# 3. Cards describe a Layout; renderers draw it

Displaying N card families in M output modes is a double-dispatch problem, and Go
has no double dispatch — so cards and renderers are decoupled by data instead.
A card exposes `Layout()` (size, a border-style enum, unstyled content, and a
semantic `Accent`) built only from gocard's own types, keeping the core free of
external dependencies; `gocard/render` defines the `Renderer` interface it
consumes and offers `render.Card(c, r)` taking `interface{ Layout() gocard.Layout }`,
while a `render/lipgloss` subpackage supplies colour for consumers who want it.
`String()` is the dependency-free floor, so `fmt.Println(card)` prints `Q♣` with
no renderer at all. There are exactly these two output paths: a one-line glyph
and a box. No third "card content" type sits between them — `Rank` and `Suit` are
already public fields.

## Considered Options

- **Go templates per card family** — the original preference, rejected because a
  template emits an opaque string that nothing downstream can measure for
  composition, template failures are per-render runtime errors inside a game loop
  where panics are forbidden, and reflection-based execution is far slower than
  direct string building.
- **`Layout` holding lipgloss types directly** — rejected because it puts
  lipgloss, termenv and `x/ansi` in every consumer's import graph including
  headless ones, and a card carrying ANSI escapes has already been rendered, so a
  plain-text renderer could not un-style it.
- **The renderer knowing concrete card types** — rejected: it cannot draw a card
  family it has never heard of.
- **The card drawing itself via renderer primitives** — rejected: the card decides
  it is drawing a box, so a plain renderer and a box-drawing renderer can no
  longer meaningfully differ.

## Consequences

- A renderer written today draws a card family invented tomorrow; a new family
  controls its own size and border without any renderer changing.
- Cards pick a border *style* and a semantic *accent*, never a literal colour —
  the theme owns the palette. A real loss of per-card control, judged worthwhile.
- `Layout` is a semver-frozen struct: added fields are cheap, existing ones are
  permanent, and renderers must tolerate any combination of them.
- Multi-card composition stays out of scope, but because `Layout` is structured
  rather than an opaque string, a future `render.Hand([]Card)` can measure cards
  and be added without a breaking change — the specific reason templates lost.
- `String()` must be on the value receiver so `Card` and `*Card` print alike.
