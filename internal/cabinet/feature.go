package cabinet

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/fastygo/framework/pkg/app"
	"github.com/fastygo/framework/pkg/web"
	"github.com/fastygo/framework/pkg/web/locale"
	"github.com/fastygo/framework/pkg/web/view"
	"github.com/fastygo/imagy/internal/auth"
	"github.com/fastygo/imagy/internal/cap"
	"github.com/fastygo/imagy/internal/fixtures"
	"github.com/fastygo/imagy/internal/paneldef"
	"github.com/fastygo/imagy/internal/views"
	"github.com/fastygo/panel"
	ui8layout "github.com/fastygo/ui8kit/layout"
)

// Feature wires cabinet HTTP routes.
type Feature struct {
	authn         auth.Authenticator
	panel         *panel.Panel[auth.Principal, cap.Capability]
	available     []string
	defaultLocale string
	navMerged     []app.NavItem
}

// NewFeature constructs the cabinet feature.
func NewFeature(authn auth.Authenticator, p *panel.Panel[auth.Principal, cap.Capability], available []string, defaultLocale string) *Feature {
	return &Feature{
		authn:         authn,
		panel:         p,
		available:     available,
		defaultLocale: defaultLocale,
	}
}

// SetNavItems implements app.NavProvider.
func (f *Feature) SetNavItems(items []app.NavItem) {
	f.navMerged = append([]app.NavItem(nil), items...)
}

// ID implements app.Feature.
func (f *Feature) ID() string {
	return "cabinet"
}

// NavItems implements app.Feature.
func (f *Feature) NavItems() []app.NavItem {
	return nil
}

func (f *Feature) panelNav() []ui8layout.NavItem {
	raw := f.panel.Registry().NavItems(paneldef.AdminPrincipal)
	out := make([]ui8layout.NavItem, len(raw))
	for i, it := range raw {
		out[i] = ui8layout.NavItem{Label: it.Label, Path: it.Path, Icon: it.Icon}
	}
	return out
}

func (f *Feature) requireAuth(w http.ResponseWriter, r *http.Request) bool {
	_, ok := f.authn.Session.Read(r)
	if ok {
		return true
	}
	q := url.Values{}
	if r.URL.Path != "" && r.URL.Path != "/cabinet/login" {
		q.Set("return_to", r.URL.Path)
	}
	loc := "/cabinet/login"
	if enc := q.Encode(); enc != "" {
		loc += "?" + enc
	}
	http.Redirect(w, r, loc, http.StatusSeeOther)
	return false
}

func (f *Feature) fixtureLocale(ctx context.Context) fixtures.Locale {
	code := locale.From(ctx)
	if code == "" {
		code = f.defaultLocale
	}
	loc, err := fixtures.LoadLocale(code)
	if err != nil {
		loc, _ = fixtures.LoadLocale(f.defaultLocale)
	}
	return loc
}

func (f *Feature) assetPaths() views.AssetPaths {
	return views.AssetPaths{
		CSS:     "/static/css/app.css",
		ThemeJS: "/static/js/theme.js",
		AppJS:   "/static/js/ui8kit.js",
	}
}

func (f *Feature) layoutData(ctx context.Context, r *http.Request, title, active string) views.LayoutData {
	fix := f.fixtureLocale(ctx)
	lt := view.BuildLanguageToggleFromContext(ctx,
		view.WithLocaleLabels(map[string]string{"en": "EN", "ru": "RU"}),
		view.WithLabel(fix.LanguageToggleLabel),
	)
	return views.LayoutData{
		Title:    title + " · " + fix.Brand,
		Lang:     locale.From(ctx),
		Brand:    fix.Brand,
		Active:   active,
		NavItems: f.panelNav(),
		Assets:   f.assetPaths(),
		Theme: ui8layout.ThemeToggleProps{
			Label:              fix.Theme.Label,
			SwitchToDarkLabel:  fix.Theme.SwitchToDarkLabel,
			SwitchToLightLabel: fix.Theme.SwitchToLight,
		},
		LanguageToggle:     lt,
		AccountEmail:       sessionEmail(r, f),
		AccountSignOutText: fix.Account.SignOut,
	}
}

func sessionEmail(r *http.Request, f *Feature) string {
	s, ok := f.authn.Session.Read(r)
	if !ok {
		return ""
	}
	return s.Email
}

// Routes implements app.Feature.
func (f *Feature) Routes(mux *http.ServeMux) {
	// Exact root only (Go 1.22+); avoids ServeMux conflict with framework's "/static/" subtree.
	mux.HandleFunc("GET /{$}", f.getRoot)
	mux.HandleFunc("GET /cabinet/login", f.getLogin)
	mux.HandleFunc("POST /cabinet/login", f.postLogin)
	mux.HandleFunc("POST /cabinet/logout", f.postLogout)
	mux.HandleFunc("GET /cabinet", f.getDashboard)
	mux.HandleFunc("GET /cabinet/sample", f.getSampleStub)
}

func (f *Feature) getRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/cabinet", http.StatusSeeOther)
}

func (f *Feature) getLogin(w http.ResponseWriter, r *http.Request) {
	if _, ok := f.authn.Session.Read(r); ok {
		http.Redirect(w, r, "/cabinet", http.StatusSeeOther)
		return
	}
	ctx := r.Context()
	fix := f.fixtureLocale(ctx)
	lt := view.BuildLanguageToggleFromContext(ctx,
		view.WithLocaleLabels(map[string]string{"en": "EN", "ru": "RU"}),
		view.WithLabel(fix.LanguageToggleLabel),
	)
	data := views.LoginPageData{
		Title:         fix.Login.Title,
		Lang:          locale.From(ctx),
		Brand:         fix.Brand,
		Subtitle:      fix.Login.Subtitle,
		Error:         "",
		ReturnTo:      r.URL.Query().Get("return_to"),
		Assets:        f.assetPaths(),
		EmailLabel:    fix.Login.EmailLabel,
		PasswordLabel: fix.Login.PasswordLabel,
		SubmitLabel:   fix.Login.SubmitLabel,
		Theme: ui8layout.ThemeToggleProps{
			Label:              fix.Theme.Label,
			SwitchToDarkLabel:  fix.Theme.SwitchToDarkLabel,
			SwitchToLightLabel: fix.Theme.SwitchToLight,
		},
		LanguageToggle: lt,
	}
	_ = web.Render(ctx, w, views.LoginPage(data))
}

func (f *Feature) postLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad form", http.StatusBadRequest)
		return
	}
	pr, ok := auth.FixtureLogin(r.PostForm.Get("email"), r.PostForm.Get("password"))
	if !ok {
		fix := f.fixtureLocale(ctx)
		lt := view.BuildLanguageToggleFromContext(ctx,
			view.WithLocaleLabels(map[string]string{"en": "EN", "ru": "RU"}),
			view.WithLabel(fix.LanguageToggleLabel),
		)
		data := views.LoginPageData{
			Title:         fix.Login.Title,
			Lang:          locale.From(ctx),
			Brand:         fix.Brand,
			Subtitle:      fix.Login.Subtitle,
			Error:         fix.Login.ErrorBadCreds,
			ReturnTo:      r.PostForm.Get("return_to"),
			Assets:        f.assetPaths(),
			EmailLabel:    fix.Login.EmailLabel,
			PasswordLabel: fix.Login.PasswordLabel,
			SubmitLabel:   fix.Login.SubmitLabel,
			Theme: ui8layout.ThemeToggleProps{
				Label:              fix.Theme.Label,
				SwitchToDarkLabel:  fix.Theme.SwitchToDarkLabel,
				SwitchToLightLabel: fix.Theme.SwitchToLight,
			},
			LanguageToggle: lt,
		}
		_ = web.Render(ctx, w, views.LoginPage(data))
		return
	}
	sess := auth.SessionFromPrincipal(pr)
	if err := f.authn.Session.Issue(w, sess); err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}
	ret := r.PostForm.Get("return_to")
	if ret == "" || !strings.HasPrefix(ret, "/cabinet") {
		ret = "/cabinet"
	}
	http.Redirect(w, r, ret, http.StatusSeeOther)
}

func (f *Feature) postLogout(w http.ResponseWriter, r *http.Request) {
	f.authn.Session.Clear(w)
	http.Redirect(w, r, "/cabinet/login", http.StatusSeeOther)
}

func (f *Feature) getDashboard(w http.ResponseWriter, r *http.Request) {
	if !f.requireAuth(w, r) {
		return
	}
	ctx := r.Context()
	fix := f.fixtureLocale(ctx)
	layout := f.layoutData(ctx, r, fix.Dashboard.Title, "/cabinet")
	_ = web.Render(ctx, w, views.CabinetLayout(layout, views.DashboardPage(views.DashboardData{
		Title:       fix.Dashboard.Title,
		Description: fix.Dashboard.Description,
		Body:        fix.Dashboard.Body,
	})))
}

func (f *Feature) getSampleStub(w http.ResponseWriter, r *http.Request) {
	if !f.requireAuth(w, r) {
		return
	}
	ctx := r.Context()
	fix := f.fixtureLocale(ctx)
	layout := f.layoutData(ctx, r, fix.SampleStub.Title, "/cabinet/sample")
	_ = web.Render(ctx, w, views.CabinetLayout(layout, views.SamplePage(views.SampleData{
		Title:       fix.SampleStub.Title,
		Description: fix.SampleStub.Description,
		Body:        fix.SampleStub.Body,
	})))
}
