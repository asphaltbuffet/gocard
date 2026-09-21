# 1. Record architecture decisions

Date: 2026-09-21

## Status

Accepted

## Context

We need to record the architectural decisions made on this project, so that
the reasoning behind them survives past the conversation that produced them.

This matters more than usual for `gocard`: it is a **public Go library**. Its
exported surface is a semver contract, so decisions about type shape (value vs.
reference), mutability, and naming are expensive to reverse once released. A
decision recorded here is cheaper to revisit than one reconstructed from a diff.

## Decision

We will use Architecture Decision Records, as described by Michael Nygard in
[Documenting Architecture Decisions][nygard].

ADRs live in `docs/adr/`, numbered sequentially and zero-padded to four digits
(`0002-…`, `0003-…`). Each record has the sections used here: Status, Context,
Decision, Consequences.

Records are immutable once Accepted. To change a decision, write a new ADR that
supersedes the old one and update the old record's Status to
`Superseded by ADR-XXXX`.

[nygard]: https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions

## Consequences

- Architectural decisions have a durable, reviewable home, and agents reading
  this repo can find the reasoning without replaying history.
- Each significant decision costs a short write-up. Decisions that are obvious
  or trivially reversible don't need one — reserve ADRs for choices that are
  expensive to undo.
- `docs/agents/domain.md` instructs skills to read this directory before
  proposing changes and to flag contradictions explicitly rather than silently
  overriding a recorded decision.
