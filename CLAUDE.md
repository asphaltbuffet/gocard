# CLAUDE.md

This file provides guidance to Claude Code when working with code in this repository.

`gocard` is a **library**. It has no binary, no CLI, and no release artifacts —
consumers get it via `go get`, and the Go module proxy resolves versions from git tags.

## VCS

Use **jujutsu** (`jj`), not `git`. Never run `git` commands directly.

## Build & Test Commands

```bash
mise run test        # gotestsum, race detector, coverage → bin/coverage.out
mise run lint        # golangci-lint → bin/golangci-lint.html
mise run lint-fix    # golangci-lint --fix
mise run generate    # go generate ./... (stringer for iota enums)
mise run dev         # generate + lint + test
mise run cover       # coverage HTML → bin/coverage.html
mise run mod-tidy    # go mod tidy

# Run a single test package
gotestsum -- -race ./internal/...

# Run one test by name
gotestsum -- -run TestFoo ./...
```

There is no `build` task — `mise run test` compiles every package. Use
`go build ./...` directly for a bare compile check.

Always run `mise run lint` and `mise run test` before committing.

## Task Completion Checklist

1. Fix all warnings (`mise run lint`, `mise run test`, compiler)
2. Run `/audit-docs` after features or fixes

## Architecture

<!-- TODO: describe the high-level architecture -->

## Repository Layout

```
gocard.go      # exported API
doc.go         # package documentation
internal/      # implementation details; not importable by consumers
CONTEXT.md     # domain glossary
docs/adr/      # architecture decision records
docs/agents/   # agent skill configuration
```

## Key Patterns

<!-- TODO: document constructor patterns, testing approach, etc. -->

## API Conventions

This is a public library — the exported surface is a contract under
[semantic versioning](https://semver.org/). Treat additions as cheap and
changes as expensive.

- **Exported identifiers need doc comments** starting with the identifier name.
  `godoclint` and `revive` enforce this.
- **Accept interfaces, return concrete types.** Keep parameter interfaces
  minimal and defined at the consumer.
- **Never export a type you are not prepared to keep.** Anything exported is
  part of the API contract; prefer `internal/` until the shape has settled.
- **Zero values should be useful** where possible, so `var d Deck` is usable
  without a constructor.
- Constructors are `New…` and return `(T, error)` only when construction can
  actually fail.
- Wrap errors with context: `fmt.Errorf("shuffling deck: %w", err)`.
  Export sentinel errors (`ErrFoo`) or typed errors when callers need to branch.
- **No panics in library code** — return errors. Panic only for programmer
  errors that cannot be signalled otherwise, and document it.
- Accept a `context.Context` as the first parameter for anything that may block.
- No type assertions / `any` casts — prefer typed interfaces.
- No magic numbers — use named constants.
- Timestamps: RFC3339 with UTC `Z` suffix.

## Testing

- `testify/assert` for non-fatal assertions, `testify/require` for preconditions
- Table-driven tests are the default
- Add testable examples (`func ExampleDeck_Shuffle`) for exported API — they
  appear in godoc and are compiled and run by `go test`
- Prefer **hand-rolled fakes** for internal interfaces; reserve `mockery`-generated
  mocks for external dependencies only

## Releasing

Releases are **tag-only**. There is no goreleaser, no GitHub Release object, and
no published binary.

1. `changie new` — describe changes
2. `mise run pre-release <major|minor|patch>` — batch + merge CHANGELOG
3. Open a PR, merge to main
4. `mise run release` — tag and push; the module proxy picks up the tag

## Agent skills

### Issue tracker

GitHub Issues on `asphaltbuffet/gocard`, via the `gh` CLI. See `docs/agents/issue-tracker.md`.

### Triage labels

Canonical vocabulary — each role's label string equals its name. See `docs/agents/triage-labels.md`.

### Domain docs

Single-context: `CONTEXT.md` + `docs/adr/` at the repo root. See `docs/agents/domain.md`.

## Shell Tooling Preferences

- `fd -H` over `find`, `rg` over `grep`, `sd` over `sed`, `jq` for JSON
- No `&&` between shell commands — run as separate tool calls
- No `git` commands — use `jj` equivalents
