package auth

import (
	"crypto/subtle"
	"strings"
)

const (
	FixtureEmail    = "test@admin.dash"
	FixturePassword = "test"
	FixtureUserID   = "fixture-admin"
)

// FixtureLogin validates identifier and password for the built-in single user.
func FixtureLogin(identifier, password string) (Principal, bool) {
	email := strings.TrimSpace(strings.ToLower(identifier))
	wantEmail := strings.ToLower(FixtureEmail)
	if len(email) != len(wantEmail) || subtle.ConstantTimeCompare([]byte(email), []byte(wantEmail)) != 1 {
		return Principal{}, false
	}
	pass := []byte(password)
	wantPass := []byte(FixturePassword)
	if len(pass) != len(wantPass) || subtle.ConstantTimeCompare(pass, wantPass) != 1 {
		return Principal{}, false
	}
	return Principal{UserID: FixtureUserID, Email: FixtureEmail}, true
}
