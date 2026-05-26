# `registry:components`

**Self-contained, reusable** templ composites — props in, markup out. No HTTP, domain fetch, or app fixtures inside the package.

## Contract

| Rule | Detail |
|------|--------|
| **Dependencies** | `github.com/fastygo/templ/ui`, `templ/components`, `templ/utils` only — plus other `internal/ui/components/*` when composed by convention |
| **Defaults** | English wireframe copy in `defaults.go` (or `placeholders.go`) so `fastygo add` / docs work without an app |
| **Overrides** | Callers and showcase pass `BlogCardData`-style structs; app `internal/fixtures` may overlay i18n at the view layer |
| **Files** | Single self-contained `*.templ` (types, defaults, helpers, markup) |

## Current packages

| Package | Role |
|---------|------|
| `blogcard/` | Vertical and horizontal blog cards with media placeholder |
| `icon/` | Latty mask icons (app shell; may move to templ later) |
| `toggles/` | Language toggle (`framework/view` — app chrome) |

## Docs

Example: `blogcard` → `/docs/components/blog-card`.

## Heuristic

Pure presentation → **`components`**. Section scaffolds → **`blocks/`** (later). Fetch/state → **`widgets/`**.
