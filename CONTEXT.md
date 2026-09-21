# Context: gocard

The domain glossary for this repo. Skills read this to learn the project's
language before proposing changes; see `docs/agents/domain.md` for the rules.

This file is **lazily filled**. Add a term when a real decision forces you to
pin its meaning down — not speculatively. An empty section is honest; a
guessed definition is worse than none.

## Domain

<!-- TODO: one paragraph — what problem does gocard model, and what does it
     deliberately NOT model? e.g. "card and deck primitives, game-agnostic;
     no game rules, no scoring, no player/table state" -->

## Glossary

Terms below are the vocabulary the exported API should use. Once a term is
exported it is part of the semver contract, so settle the name here first.

| Term | Definition | Notes |
| --- | --- | --- |
| <!-- Card --> | <!-- TODO --> | |
| <!-- Rank --> | <!-- TODO --> | |
| <!-- Suit --> | <!-- TODO --> | |
| <!-- Deck --> | <!-- TODO --> | |

### Terms to resolve

Open questions where the right word isn't settled yet. Candidates for
`/grill-with-docs`:

- **Deck vs. Shoe** — is a multi-deck shoe a distinct type or a `Deck` built
  from N standard decks?
- **Hand vs. Pile** — is an ordered face-down stack the same type as a
  player's held cards, or different?
- **Draw vs. Deal** — does drawing mutate the deck, or return a new one?
  (Related: is `Deck` a value or a reference type?)
- **Jokers** — part of a "standard" deck or an explicit opt-in?

## Deliberately avoided terms

<!-- TODO: synonyms the project rejects, and what to say instead.
     e.g. "don't say 'shuffle the pack' — it's a Deck, not a pack" -->

## See also

- `docs/adr/` — architectural decisions
- `docs/agents/domain.md` — how skills consume this file
