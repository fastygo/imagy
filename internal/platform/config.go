package platform

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/fastygo/framework/pkg/app"
)

const devSessionFallback = "dev-only-insecure-session-key-min-32bytes!"

// Config extends framework config with cabinet defaults.
type Config struct {
	app.Config
}

// Load reads environment-backed config and applies cabinet defaults.
func Load() (Config, error) {
	cfg, err := app.LoadConfig()
	if err != nil {
		return Config{}, err
	}
	// Framework's default StaticDir targets CMS-style layouts; this repo uses ./web/static.
	if strings.TrimSpace(os.Getenv("APP_STATIC_DIR")) == "" {
		cfg.StaticDir = "web/static"
	}
	if abs, err := filepath.Abs(cfg.StaticDir); err == nil {
		cfg.StaticDir = filepath.Clean(abs)
	}
	if fi, err := os.Stat(cfg.StaticDir); err != nil || !fi.IsDir() {
		slog.Warn("APP static root is missing or not a directory", "path", cfg.StaticDir, "err", err)
	}
	if strings.TrimSpace(os.Getenv("SESSION_KEY")) == "" {
		slog.Warn("SESSION_KEY is empty; using insecure development default")
		cfg.SessionKey = devSessionFallback
	}
	return Config{Config: cfg}, nil
}
