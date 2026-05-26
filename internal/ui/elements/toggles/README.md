# In-app UI registry (`internal/ui`)

Staging area for the FastyGoUI showcase. **`internal/showcase/`** holds localized Markdown and the docgen compiler; this tree holds **reusable UI**.

## Layout

```
internal/ui/
  layout/       # App shell — stays in github.com/fastygo/ui
  components/   # registry:components — templ composites (CLI-ready)
  blocks/       # registry:blocks — staging → github.com/fastygo/blocks
  widgets/      # registry:widgets — staging → github.com/fastygo/widgets
  variants/     # Optional wireframe utility maps
  utils/        # Thin helpers on github.com/fastygo/templ/utils
internal/showcase/
  content/      # Localized Markdown docs sources
  docgen/       # Static docs compiler
```

## External atoms

- `import "github.com/fastygo/templ/ui"` → `@ui.*`
- `import cmp "github.com/fastygo/templ/components"` → `@cmp.*`

## Extraction (later)

| Path | Future module |
|------|----------------|
| `blocks/*` | `github.com/fastygo/blocks` |
| `widgets/*` | `github.com/fastygo/widgets` |
| `components/*` | CLI copy / optional shared module |
| `layout/` | Stays in this app |

See `.cursor/rules/fastygo-ui-design-system-registry.mdc`.
