package layout

import (
	"fmt"
	"strings"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	uiutils "github.com/fastygo/templ/utils"
)

const (
	MobileSheetTriggerID = "ui8kit-mobile-sheet-trigger"
	MobileSheetPanelID   = "ui8kit-mobile-sheet-panel"

	NavCollapseFullVisible = 5
	NavCollapseFadeVisible = 3
)

func navCollapseThreshold() int {
	return NavCollapseFullVisible + NavCollapseFadeVisible
}

// GroupNavItems splits flat nav items into section groups (section headers start a new group).
func GroupNavItems(items []NavItem) []NavSectionGroup {
	var groups []NavSectionGroup
	var current NavSectionGroup
	for _, item := range items {
		if item.Section {
			if current.Label != "" || len(current.Items) > 0 {
				groups = append(groups, current)
			}
			current = NavSectionGroup{Label: item.Label}
			continue
		}
		current.Items = append(current.Items, item)
	}
	if current.Label != "" || len(current.Items) > 0 {
		groups = append(groups, current)
	}
	return groups
}

// NavSectionNeedsCollapse reports whether a section should use fade + expand UX.
func NavSectionNeedsCollapse(items []NavItem) bool {
	return len(items) > navCollapseThreshold()
}

// NavSectionExpanded is true when the active page is not among the always-visible links.
func NavSectionExpanded(active string, items []NavItem) bool {
	if active == "" || !NavSectionNeedsCollapse(items) {
		return false
	}
	for i, item := range items {
		if i >= NavCollapseFullVisible && item.Path == active {
			return true
		}
	}
	return false
}

// NavSectionCollapseID returns a stable DOM id for a collapsible nav section.
// Mobile and desktop sidebars render separately; pass mobile=true for the sheet nav.
func NavSectionCollapseID(label string, mobile bool) string {
	s := strings.ToLower(strings.TrimSpace(label))
	if s == "" {
		if mobile {
			return "nav-collapse-section-mobile"
		}
		return "nav-collapse-section"
	}
	var b strings.Builder
	b.WriteString("nav-collapse-")
	prevHyphen := false
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			prevHyphen = false
			continue
		}
		if !prevHyphen {
			b.WriteByte('-')
			prevHyphen = true
		}
	}
	id := strings.Trim(b.String(), "-")
	if id == "" || id == "nav-collapse" {
		id = "nav-collapse-section"
	}
	if mobile {
		return id + "-mobile"
	}
	return id
}

func navCollapseOverflowID(sectionID string) string {
	return sectionID + "-overflow"
}

func navCollapseRootStackProps(sectionID string, expanded bool) ui.StackProps {
	return ui.StackProps{
		Class: "gap-0",
		Attrs: templ.Attributes{
			"id":                sectionID,
			"data-nav-collapse": "",
			"data-nav-expanded": fmt.Sprintf("%t", expanded),
		},
	}
}

func navCollapseTeaserBoxProps(expanded bool) ui.BoxProps {
	attrs := templ.Attributes{"data-nav-collapse-teaser": ""}
	if expanded {
		attrs["hidden"] = true
	}
	return ui.BoxProps{
		Class: "nav-sidebar-teaser relative w-full",
		Attrs: attrs,
	}
}

func navCollapseTeaserHitBoxProps(sectionID string, expanded bool) ui.BoxProps {
	return ui.BoxProps{
		Class: "nav-sidebar-teaser-hit absolute inset-0 z-10 cursor-pointer",
		Attrs: templ.Attributes{
			"data-nav-collapse-expand": "",
			"aria-controls":              navCollapseOverflowID(sectionID),
			"role":                       "button",
			"tabindex":                   "0",
			"aria-expanded":              fmt.Sprintf("%t", expanded),
			"aria-label":                 "Show more navigation items",
		},
	}
}

func navCollapseTeaserItemClass(fadeIndex int) string {
	base := "nav-sidebar-teaser-item flex w-full items-center gap-2 rounded-md px-4 py-2 text-sm text-muted-foreground"
	switch fadeIndex {
	case 0:
		return uiutils.Cn(base, "nav-sidebar-teaser-item-0")
	case 1:
		return uiutils.Cn(base, "nav-sidebar-teaser-item-1")
	default:
		return uiutils.Cn(base, "nav-sidebar-teaser-item-2")
	}
}

func navCollapseOverflowStackProps(sectionID string, expanded bool) ui.StackProps {
	attrs := templ.Attributes{
		"id":                         navCollapseOverflowID(sectionID),
		"data-nav-collapse-overflow": "",
	}
	if !expanded {
		attrs["hidden"] = true
	}
	return ui.StackProps{
		Class: "nav-sidebar-overflow gap-0",
		Attrs: attrs,
	}
}

func shellHeaderTitle(props ShellProps) string {
	if strings.TrimSpace(props.HeaderTitle) != "" {
		return props.HeaderTitle
	}
	return props.Title
}

func shellBrand(name string) string {
	if name == "" {
		return "App"
	}
	return name
}

func shellLang(value string) string {
	if value == "" {
		return "en"
	}
	return value
}

func shellBodyClass(props ShellProps) string {
	base := "min-h-screen overflow-x-hidden bg-background font-sans text-foreground"
	if props.MarketingShell {
		return base
	}
	return uiutils.Cn(base, "max-md:has-[#ui8kit-mobile-sheet-panel:not([hidden])]:overflow-hidden")
}

func shellHasNavigation(props ShellProps) bool {
	return !props.MarketingShell && len(props.NavItems) > 0
}

func isExternalNavLink(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}

func themeToggleLabel(value string) string {
	if value == "" {
		return "Toggle theme"
	}
	return value
}

func themeToggleSwitchToDarkLabel(value string) string {
	if value == "" {
		return "Switch to dark theme"
	}
	return value
}

func themeToggleSwitchToLightLabel(value string) string {
	if value == "" {
		return "Switch to light theme"
	}
	return value
}

func sidebarItemClasses(active, path string) string {
	base := "flex w-full items-center gap-2 rounded-md px-4 py-2 text-sm"
	if active == path {
		return uiutils.Cn(base, "bg-accent text-accent-foreground")
	}
	return uiutils.Cn(base, "text-muted-foreground hover:bg-accent")
}

func headerMenuButtonProps() ui.ButtonProps {
	return ui.ButtonProps{
		ID:      MobileSheetTriggerID,
		Type:    "button",
		Variant: "unstyled",
		Class: uiutils.Cn(
			"inline-flex h-8 w-8 shrink-0 items-center justify-center rounded-md border border-border text-foreground md:hidden",
		),
		Attrs: uiutils.MergeAttrs(
			templ.Attributes{
				"data-ui8kit-dialog-open":   true,
				"data-ui8kit-dialog-target": MobileSheetPanelID,
			},
			uiutils.AriaLabel("Open navigation menu"),
			uiutils.AriaHasPopup("dialog"),
			uiutils.AriaControls(MobileSheetPanelID),
			uiutils.AriaExpanded(false),
		),
	}
}

func themeToggleButtonProps(props ThemeToggleProps) ui.ButtonProps {
	return ui.ButtonProps{
		ID:      "ui8kit-theme-toggle",
		Type:    "button",
		Variant: "unstyled",
		Class: uiutils.Cn(
			"inline-flex h-8 w-8 items-center justify-center rounded-md bg-transparent p-0 text-muted-foreground transition-colors hover:bg-accent hover:text-accent-foreground",
		),
		Attrs: uiutils.MergeAttrs(
			templ.Attributes{
				"data-switch-to-dark-label":  themeToggleSwitchToDarkLabel(props.SwitchToDarkLabel),
				"data-switch-to-light-label": themeToggleSwitchToLightLabel(props.SwitchToLightLabel),
				"title":                      themeToggleLabel(props.Label),
			},
			uiutils.AriaLabel(themeToggleLabel(props.Label)),
			uiutils.AriaPressed(false),
		),
	}
}

func mobileSheetRootBlock() ui.BlockProps {
	return ui.BlockProps{
		ID:    MobileSheetPanelID,
		Class: "fixed inset-y-0 left-0 z-50 w-full md:hidden",
		Attrs: templ.Attributes{
			"data-ui8kit":        "sheet",
			"data-ui8kit-dialog": true,
			"role":               "dialog",
			"aria-modal":         "true",
			"aria-label":         "Navigation menu",
			"aria-labelledby":    "ui8kit-mobile-sheet-title",
			"data-state":         "closed",
			"hidden":             true,
		},
	}
}

func mobileSheetOverlayBox() ui.BoxProps {
	return ui.BoxProps{
		Class: "absolute inset-0 cursor-pointer bg-card/50",
		Attrs: templ.Attributes{
			"data-ui8kit-dialog-overlay": true,
			"data-ui8kit-dialog-close":   true,
			"data-ui8kit-dialog-target":  MobileSheetPanelID,
		},
	}
}

func mobileSheetCloseButtonProps() ui.ButtonProps {
	return ui.ButtonProps{
		Type:      "button",
		Variant:   "unstyled",
		Class:     "inline-flex h-8 w-8 cursor-pointer items-center justify-center rounded-md border border-border text-muted-foreground",
		AriaLabel: "Close navigation menu",
		Attrs: templ.Attributes{
			"data-ui8kit-dialog-close":  true,
			"data-ui8kit-dialog-target": MobileSheetPanelID,
		},
	}
}

func sidebarLinkButtonProps(active string, item NavItem, extraClass string) ui.ButtonProps {
	attrs := templ.Attributes{}
	if active == item.Path {
		attrs = uiutils.MergeAttrs(attrs, uiutils.AriaCurrent("page"))
	}
	if isExternalNavLink(item.Path) {
		attrs["target"] = "_blank"
		attrs["rel"] = "noopener noreferrer"
	}
	className := sidebarItemClasses(active, item.Path)
	if extraClass != "" {
		className = uiutils.Cn(className, extraClass)
	}
	return ui.ButtonProps{
		Href:    item.Path,
		Variant: "unstyled",
		Class:   className,
		Attrs:   attrs,
	}
}
