package auth_test

import (
	"testing"

	"github.com/fastygo/imagy/internal/auth"
)

func TestFixtureLogin(t *testing.T) {
	p, ok := auth.FixtureLogin("test@admin.dash", "test")
	if !ok || p.UserID != auth.FixtureUserID {
		t.Fatalf("expected fixture admin, got ok=%v principal=%+v", ok, p)
	}
	if _, ok := auth.FixtureLogin("wrong@example.test", "admin"); ok {
		t.Fatal("expected failure for wrong email")
	}
	if _, ok := auth.FixtureLogin("test@admin.dash", "nope"); ok {
		t.Fatal("expected failure for wrong password")
	}
}
