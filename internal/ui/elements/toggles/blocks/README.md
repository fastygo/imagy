# `registry:blocks`

Section-level wireframe organisms (hero, feature grid, dashboard sections, …). **shadcn-style blocks**: self-contained templ + **default English copy** in-package (`defaults.go` / `placeholders.go`).

## Groups

| Package | Showcase focus |
|---------|----------------|
| `dashboard/` | Dashboard / home sections |
| `marketing/` | Landing, hero, CTA scaffolds |
| `docs/` | Documentation-style sections |

Add new top-level folders only when a **new showcase group** is needed (e.g. `storefront/`, `editorial/`).

## Rules

- Compose with `github.com/fastygo/templ/ui` and `templ/components` — **no raw HTML tags**.
- Tailwind + semantic tokens only; pass **ui8px** policy.
- **Do not** import `github.com/fastygo/blocks` during active development in this repo.
- Interactive patterns: correct `data-ui8kit` + static ARIA + manifest (see `fastygo-ui-aria.mdc`).

## Extraction

When a block is stable, move the package to **`github.com/fastygo/blocks/<name>`** and `require` it from this app. Keep default data inside the block package; app `internal/fixtures` only for i18n overlay.
