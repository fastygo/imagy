package layout

import "github.com/a-h/templ"

// NavItem describes a single sidebar navigation link.
// When Section is true, Label is a non-interactive group heading (Path is ignored).
type NavItem struct {
	Path    string
	Label   string
	Icon    string
	Section bool
}

// NavSectionGroup is a sidebar section heading plus its link items.
type NavSectionGroup struct {
	Label string
	Items []NavItem
}

// SidebarProps configures the sidebar navigation.
type SidebarProps struct {
	Items  []NavItem
	Active string
	Mobile bool
}

// HeaderProps configures the top header bar.
type HeaderProps struct {
	ShowMenuTrigger      bool
	Title                string
	Extra                templ.Component
	Trailing             templ.Component
	ThemeToggle          ThemeToggleProps
	ThemeToggleComponent templ.Component
}

// ThemeToggleProps configures copy for the theme toggle button.
type ThemeToggleProps struct {
	Label              string
	SwitchToDarkLabel  string
	SwitchToLightLabel string
}

// ShellProps configures the full page shell (sidebar + header + main).
type ShellProps struct {
	Title                string // document title for <title> (SEO)
	HeaderTitle          string // visible page title in the header bar
	Lang                 string
	BrandName            string
	Active               string
	NavItems             []NavItem
	HeadExtra            templ.Component
	HeaderExtra          templ.Component
	HeaderTrailing       templ.Component
	ThemeToggle          ThemeToggleProps
	ThemeToggleComponent templ.Component
	MarketingShell       bool
}
