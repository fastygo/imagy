# `registry:variants`

Optional **Go maps of allowed utility strings** for wireframe variants (button density, card padding, …). Use when the same component needs named, policy-safe class presets without hiding structure in opaque helpers.

## Example shape

```go
var ButtonDensity = map[string]string{
    "compact": "h-8 px-3 text-sm",
    "default": "h-10 px-4 py-2",
}
```

Every token must pass **`bun run lint:ui8px`**.
