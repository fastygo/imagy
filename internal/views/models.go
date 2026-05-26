package views

import (
	"github.com/fastygo/framework/pkg/web/view"
	ui8layout "github.com/fastygo/ui8kit/layout"
)

// AssetPaths are URLs for CSS and JS bundles.
type AssetPaths struct {
	CSS     string
	ThemeJS string
	AppJS   string
}

// LayoutData drives the cabinet shell (sidebar + header).
type LayoutData struct {
	Title              string
	Lang               string
	Brand              string
	Active             string
	NavItems           []ui8layout.NavItem
	Assets             AssetPaths
	Theme              ui8layout.ThemeToggleProps
	LanguageToggle     view.LanguageToggleData
	AccountEmail       string
	AccountSignOutText string
}

// LoginPageData is the sign-in screen.
type LoginPageData struct {
	Title          string
	Lang           string
	Brand          string
	Subtitle       string
	Error          string
	ReturnTo       string
	Assets         AssetPaths
	EmailLabel     string
	PasswordLabel  string
	SubmitLabel    string
	Theme          ui8layout.ThemeToggleProps
	LanguageToggle view.LanguageToggleData
}

// DashboardData is the cabinet home page body.
type DashboardData struct {
	Title       string
	Description string
	Body        string
}

// SampleData is a second stub route for copy-paste onboarding.
type SampleData struct {
	Title       string
	Description string
	Body        string
}
