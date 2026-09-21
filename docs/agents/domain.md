# Domain Docs

How the engineering skills should consume this repo's domain documentation when exploring the codebase.

This repo is **single-context**: one `CONTEXT.md` and one `docs/adr/` at the root.

## Before exploring, read these

- **`CONTEXT.md`** at the repo root — the domain glossary.
- **`docs/adr/`** — read ADRs that touch the area you're about to work in.

There is no `CONTEXT-MAP.md` and no per-context `docs/adr/`; if you find
yourself looking for one, the repo layout has changed and this file is stale.

If any of these files don't exist, **proceed silently**. Don't flag their absence; don't suggest creating them upfront. The producer skill (`/grill-with-docs`) creates them lazily when terms or decisions actually get resolved.

## File structure

```
/
├── CONTEXT.md                 ← domain glossary
├── docs/
│   ├── adr/                   ← architecture decision records
│   └── agents/                ← this directory (skill configuration)
├── gocard.go                  ← exported API
├── doc.go                     ← package documentation
└── internal/                  ← implementation details
```

Note this is a single-module Go library with no `src/` directory — the package
lives at the repo root. ADRs are always at `docs/adr/`, never nested.

## Use the glossary's vocabulary

When your output names a domain concept (in an issue title, a refactor proposal, a hypothesis, a test name), use the term as defined in `CONTEXT.md`. Don't drift to synonyms the glossary explicitly avoids.

For a playing-card library the glossary matters more than usual: terms like
*card*, *rank*, *suit*, *deck*, *shoe*, *hand*, *pile*, *draw*, *deal* and
*shuffle* have precise, non-interchangeable meanings that vary between card
games. Pin them down in `CONTEXT.md` before the API hardens — once exported,
a name is part of the semver contract.

If the concept you need isn't in the glossary yet, that's a signal — either you're inventing language the project doesn't use (reconsider) or there's a real gap (note it for `/grill-with-docs`).

## Flag ADR conflicts

If your output contradicts an existing ADR, surface it explicitly rather than silently overriding:

> _Contradicts ADR-0003 (immutable Deck values) — but worth reopening because…_
