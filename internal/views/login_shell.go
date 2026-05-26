package views

import (
	etoggles "github.com/fastygo/imagy/internal/ui/elements/toggles"
	layout "github.com/fastygo/ui8kit/layout"
)

// LoginShellProps builds UI8Kit shell props for the marketing-style login screen.
func LoginShellProps(d LoginPageData) layout.ShellProps {
	return layout.ShellProps{
		Title:          "Dash · " + d.Title + " · " + d.Brand,
		Lang:           d.Lang,
		BrandName:      d.Brand,
		Active:         "",
		NavItems:       nil,
		CSSPath:        d.Assets.CSS,
		ThemeJSPath:    d.Assets.ThemeJS,
		AppJSPath:      d.Assets.AppJS,
		HeaderExtra:    etoggles.LanguageToggle(d.LanguageToggle),
		ThemeToggle:    d.Theme,
		MarketingShell: true,
	}
}
