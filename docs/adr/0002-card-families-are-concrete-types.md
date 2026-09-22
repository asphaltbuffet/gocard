# 2. Card families are concrete types; Deck is generic over them

Each card family is its own concrete struct — `gocard.Card` is a French-suited
playing card, and Tarot or Magic-style cards would be separate types — so
comparing cards from different families is a compile error rather than a silent
`false`. Because a Deck therefore holds one family, `Deck[T]` is generic, letting
the shuffle and draw machinery be written once and reused by any card type,
including third-party ones; `Standard` aliases `Deck[Card]`. A Joker is a `Card`
with rank `Joker` and no suit, which keeps a joker-bearing deck homogeneous.

## Considered Options

- **`Card` as an interface** — rejected: unrelated card types would compare equal
  to `false` instead of failing to compile, and an exported interface cannot gain
  methods without breaking implementors, so it would be frozen before the
  rendering design settled.
- **One concrete `Card` struct for every family**, with optional `Rank`, `Suit`
  and `Label` — rejected: keeps `Deck` non-generic and reads better in godoc, but
  re-introduces the equality problem and leaves most fields meaningless for most
  values.

## Consequences

- `var d gocard.Deck` does not compile; a generic type needs its parameter. This
  costs the project's "zero values should be useful" aim, accepted because
  writing Fisher–Yates once beats once per family.
- The zero values of `Rank` and `Suit` must mean **absent**, so their `iota`
  sequences need an explicit "none" member — otherwise every suitless card
  silently becomes a club.
- `Card` stays comparable with `==` and usable as a map key.
