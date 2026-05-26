# Imagy Panel

Private single-user **cabinet** built on [FastyGo Framework](https://github.com/fastygo/framework), [Panel](https://github.com/fastygo/panel), and [UI8Kit](https://github.com/fastygo/ui8kit) `v0.4.0`. This repository is a **cloneable shell**: admin layout (sidebar, mobile sheet, header, language and theme toggles), fixture authentication, and embedded locale JSON. Use it as a neutral starting point for a new cabinet.

## Prerequisites

- Go 1.25+
- [Bun](https://bun.sh) (for CSS build and `ui8px`)

## Quick start

```bash
bun install
go mod download
bun run build:css
go tool templ generate ./...
export SESSION_KEY="replace-with-a-long-random-secret-at-least-32-chars"
bun run go
```

UI8Kit static CSS/JS, theme scripts, and **Google Sans** (`gfonts.css` + `fonts/google-sans/`) are committed under [`web/static/`](web/static/).

`bun run go` runs [`scripts/run-server.mjs`](scripts/run-server.mjs): the server always starts with the **repository root as cwd** (correct `web/static`), and **Ctrl+C** is forwarded to the Go process so the port is released.

Closing a **browser tab does not** stop an HTTP server. Stop the job in the terminal (**Ctrl+C**) or close the terminal panel; if the port stays busy, an old `go` process is still running (see troubleshooting below).

Open [http://127.0.0.1:8080/](http://127.0.0.1:8080/) — you are redirected to `/cabinet` (or to login if not signed in). You can also open [http://127.0.0.1:8080/cabinet/login](http://127.0.0.1:8080/cabinet/login) directly.

### Default operator (fixture)

- **Email:** `test@admin.dash`
- **Password:** `test`

There is **no SQLite** and no server-side user database; credentials are fixed in code for this autonomous template.

## Environment

| Variable | Default | Purpose |
|----------|---------|---------|
| `APP_BIND` | `127.0.0.1:8080` | HTTP listen address |
| `APP_STATIC_DIR` | `web/static` when env omitted | Static files under `/static/`. Framework’s built-in default points at a CMS-style folder; this app **forces** `web/static` whenever `APP_STATIC_DIR` is not set in the environment. Use an absolute path if you do not start the server from the repo root. |
| `SESSION_KEY` | dev-only fallback (logged) | HMAC secret for the session cookie |
| `APP_DEFAULT_LOCALE` | `en` | Default locale |
| `APP_AVAILABLE_LOCALES` | `en,ru` | Locales for the header switcher (query + cookie) |

Probes: `GET /healthz` and `GET /readyz` are registered in [`cmd/server/main.go`](cmd/server/main.go).

Set `SESSION_KEY` to a long random value before any non-local use.

## Troubleshooting

### `listen tcp ... bind: Only one usage of each socket address`

Another process (often a previous `go run`) is still bound to that port. Stop it or use another port:

```bash
export APP_BIND=127.0.0.1:8081
bun run go
```

On Windows, find and end the listener, for example: `netstat -ano | findstr :8080` then `taskkill /PID <pid> /F`.

### Static files 404

Run from the repo root (or use `bun run go`), run `bun run build:css`, and ensure `web/static` exists. See `APP_STATIC_DIR` in the table above.

## Project layout

| Path | Role |
|------|------|
| [`cmd/server/main.go`](cmd/server/main.go) | Composition root: config, locales, health, cabinet feature |
| [`internal/cabinet/`](internal/cabinet/) | HTTP routes: `/cabinet`, `/cabinet/login`, `/cabinet/logout`, `/cabinet/sample` (placeholder) |
| [`internal/auth/`](internal/auth/) | Cookie session + fixture login |
| [`internal/paneldef/`](internal/paneldef/) | `panel.Panel` descriptor (pages, nav metadata) |
| [`internal/fixtures/locale/`](internal/fixtures/locale/) | Embedded JSON copy per locale |
| [`internal/views/`](internal/views/) | `templ` pages, [`layout.templ`](internal/views/layout.templ) (`CabinetLayout` + UI8Kit `Shell`), [`partials/account_menu.templ`](internal/views/partials/account_menu.templ), and [`login_shell.go`](internal/views/login_shell.go) for marketing login shell |
| [`internal/ui/elements/`](internal/ui/elements/) | Small reusable UI (e.g. `elements/toggles` language control) |
| [`web/static/`](web/static/) | `app.css` (Tailwind build), `css/ui8kit/*`, `css/gfonts.css`, `fonts/google-sans/*`, `js/theme.js`, `js/ui8kit.js` |

## Verification

```bash
bun run verify
```

This runs `templ generate`, Tailwind build, `ui8px lint` (policy under [`.ui8px/policy/`](.ui8px/policy/)), and `go test ./...`.

## Cloning this template for a new cabinet

1. Copy the repository (or subtree: `cmd/server`, `internal/cabinet`, `internal/ui`, `internal/views`, `internal/paneldef`, `web/static`, `package.json`, `.ui8px`).
2. Change the Go module path in `go.mod` and imports.
3. Extend [`internal/paneldef/panel.go`](internal/paneldef/panel.go) with new `panel.Page` or `panel.Resource` entries; mirror each with a route handler in [`internal/cabinet/feature.go`](internal/cabinet/feature.go).
4. Add or extend `templ` under [`internal/views/`](internal/views/) (and Tailwind `@source` in [`web/static/css/input.css`](web/static/css/input.css) for new paths).
5. Replace fixture auth in [`internal/auth/fixture.go`](internal/auth/fixture.go) with your own policy when you outgrow the single-user model.

## Roadmap

- Domain-specific pages and Panel resources wired to your product.
- Optional persistence or APIs as needed for your cabinet.

The [`.fastygo/`](.fastygo/) directory in some workspaces is reference-only and is **not** imported at runtime; this module builds standalone.
