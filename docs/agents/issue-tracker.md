# Issue tracker: GitHub

Issues and PRDs for this repo live as GitHub issues on
[`asphaltbuffet/gocard`](https://github.com/asphaltbuffet/gocard).
Use the `gh` CLI for all operations.

## Conventions

- **Create an issue**: `gh issue create --title "..." --body "..."`. Use a heredoc for multi-line bodies.
- **Read an issue**: `gh issue view <number> --comments`, filtering comments by `jq` and also fetching labels.
- **List issues**: `gh issue list --state open --json number,title,body,labels,comments --jq '[.[] | {number, title, body, labels: [.labels[].name], comments: [.comments[].body]}]'` with appropriate `--label` and `--state` filters.
- **Comment on an issue**: `gh issue comment <number> --body "..."`
- **Apply / remove labels**: `gh issue edit <number> --add-label "..."` / `--remove-label "..."`
- **Close**: `gh issue close <number> --comment "..."`

Infer the repo from `git remote -v` — `gh` does this automatically when run inside a clone.

> Note: this repo uses **jujutsu** (`jj`) for version control, with a colocated
> `.git` directory. `gh` reads that `.git` and works normally. Use `jj` for all
> history operations; `gh` is only for issue/PR interaction.

## When a skill says "publish to the issue tracker"

Create a GitHub issue.

## When a skill says "fetch the relevant ticket"

Run `gh issue view <number> --comments`.

## Labels available on this repo

Triage labels are documented separately in [`triage-labels.md`](./triage-labels.md).
Beyond those, the repo carries:

| Label | Meaning |
| --- | --- |
| `bug` | Something isn't working |
| `documentation` | Improvements or additions to documentation |
| `enhancement` | New feature or request |
| `question` | Further information is requested |
| `duplicate` | This issue or pull request already exists |
| `invalid` | This doesn't seem right |
| `good first issue` | Good for newcomers |
| `help wanted` | Extra attention is needed |
| `chore` | Maintenance task, no user-facing change |
| `refactor` | Code restructuring without behavior change |
| `test` | Test additions or fixes |
| `ci` | CI/CD pipeline changes |
| `deps` | Dependency updates |

Create a new label with
`gh label create "<name>" --color "<hex>" --description "<text>"`.
