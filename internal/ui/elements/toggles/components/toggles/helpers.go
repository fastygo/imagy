package toggles

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/fastygo/framework/pkg/web/view"
)

func languageToggleLabel(label, locale string) string {
	if strings.TrimSpace(label) != "" {
		return label
	}
	return strings.ToUpper(strings.TrimSpace(locale))
}

func languageToggleAriaLabel(data view.LanguageToggleData) string {
	if strings.TrimSpace(data.Label) != "" {
		return data.Label
	}
	return "Switch language"
}

func languageToggleAttrs(data view.LanguageToggleData) templ.Attributes {
	attrs := templ.Attributes{
		"data-default-locale":    data.DefaultLocale,
		"data-current-locale":    data.CurrentLocale,
		"data-next-locale":       data.NextLocale,
		"data-next-label":        data.NextLabel,
		"data-available-locales": strings.Join(data.AvailableLocales, ","),
	}
	if data.EnhanceWithJS {
		attrs["data-ui8kit-spa-lang"] = "1"
		attrs["data-spa-target"] = data.SPATarget
	}
	return attrs
}
