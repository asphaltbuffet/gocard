# Triage Labels

The skills speak in terms of five canonical triage roles. This file maps those roles to the actual label strings used in this repo's issue tracker.

This repo uses the **canonical vocabulary** — each role's label string equals its
name, so the mapping is an identity.

| Label in skills            | Label in our tracker | Meaning                                  |
| -------------------------- | -------------------- | ---------------------------------------- |
| `needs-triage`             | `needs-triage`       | Maintainer needs to evaluate this issue  |
| `needs-info`               | `needs-info`         | Waiting on reporter for more information |
| `ready-for-agent`          | `ready-for-agent`    | Fully specified, ready for an AFK agent  |
| `ready-for-human`          | `ready-for-human`    | Requires human implementation            |
| `wontfix`                  | `wontfix`            | Will not be actioned                     |

When a skill mentions a role (e.g. "apply the AFK-ready triage label"), use the corresponding label string from this table.

Edit the right-hand column to match whatever vocabulary you actually use.

## History

The repo was originally scaffolded with `to_triage` and `ready_afk` (from the
`/new-go-project` label set). Those were **replaced** by the canonical
`needs-triage` and `ready-for-agent` and deleted from the repo, so there is only
one label per role. If you see `to_triage` or `ready_afk` referenced anywhere,
it is stale.
