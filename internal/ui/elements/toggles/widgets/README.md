# `registry:widgets`

**UI + behavior** — loading state, user actions, API calls, timers, glue to app services. Composes `components`, `blocks`, and `github.com/fastygo/templ/*` internally.

## Rules

- Handlers / features pass **resolved** view models; widgets may call services but **must not** own routing.
- Prefer composing existing **blocks** before duplicating section markup.
- Same styling contract as the rest of the registry: Tailwind + tokens, ui8px, no raw tags.

## Extraction

When stable → **`github.com/fastygo/widgets/<name>`**. App keeps routes and wires dependencies.
