package fixtures

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed locale/*.json
var localeFS embed.FS

// Locale holds embedded copy for one language.
type Locale struct {
	Brand string `json:"brand"`
	Login struct {
		Title         string `json:"title"`
		Subtitle      string `json:"subtitle"`
		EmailLabel    string `json:"email_label"`
		PasswordLabel string `json:"password_label"`
		SubmitLabel   string `json:"submit_label"`
		ErrorBadCreds string `json:"error_bad_creds"`
	} `json:"login"`
	Theme struct {
		Label             string `json:"label"`
		SwitchToDarkLabel string `json:"switch_to_dark"`
		SwitchToLight     string `json:"switch_to_light"`
	} `json:"theme"`
	LanguageToggleLabel string `json:"language_toggle_label"`
	Account             struct {
		SignOut string `json:"sign_out"`
	} `json:"account"`
	Dashboard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"dashboard"`
	SampleStub struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"sample_stub"`
}

// LoadLocale reads locale/{code}.json (e.g. en, ru).
func LoadLocale(code string) (Locale, error) {
	raw, err := localeFS.ReadFile("locale/" + code + ".json")
	if err != nil {
		return Locale{}, fmt.Errorf("fixtures: read locale %q: %w", code, err)
	}
	var out Locale
	if err := json.Unmarshal(raw, &out); err != nil {
		return Locale{}, fmt.Errorf("fixtures: parse locale %q: %w", code, err)
	}
	return out, nil
}
