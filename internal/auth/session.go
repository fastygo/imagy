package auth

import (
	"net/http"
	"time"

	"github.com/fastygo/framework/pkg/auth"
	"github.com/fastygo/imagy/internal/cap"
)

// SessionData is stored in the signed session cookie.
type SessionData struct {
	UserID       string   `json:"user_id"`
	Email        string   `json:"email"`
	Capabilities []string `json:"capabilities"`
	IssuedAtUnix int64    `json:"issued_at_unix,omitempty"`
}

// Authenticator issues and reads browser sessions.
type Authenticator struct {
	Session auth.CookieSession[SessionData]
}

// Options configures cookie names and TTLs.
type Options struct {
	SessionName     string
	SessionPath     string
	SessionTTL      time.Duration
	SessionSecure   bool
	SessionSameSite http.SameSite
}

// NewAuthenticator builds a cookie session authenticator.
func NewAuthenticator(sessionSecret string, opts Options) Authenticator {
	if opts.SessionName == "" {
		opts.SessionName = "imagy_session"
	}
	if opts.SessionPath == "" {
		opts.SessionPath = "/"
	}
	if opts.SessionTTL <= 0 {
		opts.SessionTTL = 12 * time.Hour
	}
	if opts.SessionSameSite == 0 {
		opts.SessionSameSite = http.SameSiteLaxMode
	}
	return Authenticator{
		Session: auth.CookieSession[SessionData]{
			Name:     opts.SessionName,
			Path:     opts.SessionPath,
			Secret:   sessionSecret,
			TTL:      opts.SessionTTL,
			HTTPOnly: true,
			Secure:   opts.SessionSecure,
			SameSite: opts.SessionSameSite,
		},
	}
}

// PrincipalFromSession maps session payload to a panel principal.
func PrincipalFromSession(s SessionData) Principal {
	return Principal{UserID: s.UserID, Email: s.Email}
}

// SessionFromPrincipal builds a session for Issue().
func SessionFromPrincipal(p Principal) SessionData {
	return SessionData{
		UserID:       p.UserID,
		Email:        p.Email,
		Capabilities: []string{string(cap.Admin)},
		IssuedAtUnix: time.Now().UTC().Unix(),
	}
}
