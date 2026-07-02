# Contributing

yllmcode is early-stage software. Keep changes focused, easy to review, and
aligned with the current implementation phase.

## Development Workflow

- Work in small commits that each explain one coherent change.
- Commit after finishing each meaningful unit of work instead of batching a
  large mixed change at the end.
- Keep generated artifacts, local handoff notes, and editor state out of the
  repository.
- Run `go test ./...` before opening a pull request or merging local work.
- Avoid adding dependencies until the standard library is clearly insufficient.

## Commit Messages

Write commit messages for a future maintainer reading history without the
current conversation in front of them.

Use a short imperative subject:

```text
Add project initializer templates
Document bootstrap CLI scope
Preserve existing planning files
```

When the change needs context, add a body that explains why the change exists
and what tradeoff it makes. Avoid generic subjects such as `update files`,
`misc fixes`, or `AI changes`.

## Code Comments

Comments should explain intent, constraints, or non-obvious behavior. Do not
comment on mechanics that are already clear from the code.

Prefer comments like:

```go
// Preserve user-authored planning files unless the caller explicitly opts in.
```

Avoid comments like:

```go
// Set force to true.
```

Do not add boilerplate comments that describe who or what generated a change.
The code and commit history should stand on their own technical merits.

## Review Style

Review comments should be specific and actionable. Point to the behavior, risk,
or missing test, and suggest the smallest practical fix.

Good review comments usually answer:

- What can go wrong?
- Where does it happen?
- What would make the change easier to trust?

Avoid performative language, vague approval, and broad criticism that does not
map to a concrete change.

## Project Scope

The bootstrap phase is limited to the CLI skeleton and `.yllmcode`
initialization. Do not add planning chat, daemon behavior, model-provider logic,
internet access, or autonomous coding loops unless that phase has been approved.
