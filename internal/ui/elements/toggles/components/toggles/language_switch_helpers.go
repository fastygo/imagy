package toggles

import (
	"strings"

	"github.com/a-h/templ"
)

func languageSwitchAttrs(data LanguageSwitchProps) templ.Attributes {
	attrs := templ.Attributes{}
	if strings.TrimSpace(data.AriaLabel) != "" {
		attrs["aria-label"] = strings.TrimSpace(data.AriaLabel)
	}
	attrs["role"] = "group"
	return attrs
}

func languageSwitchItemAriaLabel(groupLabel string, item LanguageSwitchItem) string {
	if strings.TrimSpace(groupLabel) != "" {
		return groupLabel + ": " + item.Label
	}
	return "Switch to " + item.Label
}
