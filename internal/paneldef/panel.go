package paneldef

import (
	"fmt"

	"github.com/fastygo/imagy/internal/auth"
	"github.com/fastygo/imagy/internal/cap"
	"github.com/fastygo/panel"
)

// AdminPrincipal is used when aggregating static navigation for the single operator.
var AdminPrincipal = auth.Principal{UserID: auth.FixtureUserID, Email: auth.FixtureEmail}

// BuildPanel constructs the cabinet control-plane descriptor (navigation + schemas).
func BuildPanel() (*panel.Panel[auth.Principal, cap.Capability], error) {
	p, err := panel.NewPanel[auth.Principal, cap.Capability](panel.PanelOptions[cap.Capability]{
		ID:       "cabinet",
		Title:    "Cabinet",
		BasePath: "/cabinet",
	})
	if err != nil {
		return nil, err
	}
	if err := p.AddPages(
		panel.Page[cap.Capability]{
			ID:          "dashboard",
			Kind:        panel.PageDashboard,
			Title:       "Dashboard",
			Description: "Cabinet overview",
			Path:        "/cabinet",
			Icon:        "home",
			Navigation: panel.MenuItem[cap.Capability]{
				ID:    "dashboard",
				Label: "Dashboard",
				Path:  "/cabinet",
				Icon:  "home",
				Order: 0,
			},
			Capability: cap.Admin,
		},
		panel.Page[cap.Capability]{
			ID:          "sample",
			Kind:        panel.PageCustom,
			Title:       "Sample",
			Description: "Placeholder second page",
			Path:        "/cabinet/sample",
			Icon:        "layout-dashboard",
			Navigation: panel.MenuItem[cap.Capability]{
				ID:    "sample",
				Label: "Sample",
				Path:  "/cabinet/sample",
				Icon:  "layout-dashboard",
				Order: 10,
			},
			Capability: cap.Admin,
		},
	); err != nil {
		return nil, fmt.Errorf("add pages: %w", err)
	}
	return p, nil
}
