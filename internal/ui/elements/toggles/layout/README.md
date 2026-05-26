# `registry:layout`

Application-owned structural shell — **stays in `github.com/fastygo/ui`** after blocks/widgets are extracted.

- **Shell** — document, mobile sheet nav, desktop sidebar, main slot
- **Sidebar** — nav links with latty icons (`internal/ui/components/icon`)
- **Header** — menu trigger, title, theme toggle, trailing slot (language)

Uses `github.com/fastygo/templ/ui` for primitives and preserves `data-ui8kit-*` hooks for `theme.js` / `ui8kit.js` / `@ui8kit/aria`.
