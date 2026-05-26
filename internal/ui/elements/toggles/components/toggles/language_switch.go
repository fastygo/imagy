package toggles

// LanguageSwitchItem is one locale link in the header switcher.
type LanguageSwitchItem struct {
	Locale string
	Label  string
	Href   string
	Active bool
}

// LanguageSwitchProps drives the compact En / Ru header control.
type LanguageSwitchProps struct {
	AriaLabel string
	Items     []LanguageSwitchItem
}
