package auth

import "github.com/fastygo/imagy/internal/cap"

// Principal is the signed-in operator (single fixture user in v1).
type Principal struct {
	UserID string
	Email  string
}

// Has reports whether the principal holds a capability.
func (p Principal) Has(c cap.Capability) bool {
	if p.UserID == "" {
		return false
	}
	return c == cap.Admin
}
