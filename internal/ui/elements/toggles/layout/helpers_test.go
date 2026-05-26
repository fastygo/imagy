package layout

import (
	"fmt"
	"testing"
)

func TestGroupNavItems_splitsSectionHeaders(t *testing.T) {
	items := []NavItem{
		{Label: "Home", Path: "/"},
		{Section: true, Label: "Docs"},
		{Label: "Intro", Path: "/docs/introduction/"},
		{Label: "Install", Path: "/docs/installation/"},
	}
	groups := GroupNavItems(items)
	if len(groups) != 2 {
		t.Fatalf("groups: got %d want 2", len(groups))
	}
	if groups[0].Label != "" || len(groups[0].Items) != 1 {
		t.Fatalf("first group: %+v", groups[0])
	}
	if groups[1].Label != "Docs" || len(groups[1].Items) != 2 {
		t.Fatalf("second group: %+v", groups[1])
	}
}

func TestNavSectionNeedsCollapse(t *testing.T) {
	short := make([]NavItem, navCollapseThreshold())
	if NavSectionNeedsCollapse(short) {
		t.Fatal("expected no collapse at threshold")
	}
	long := make([]NavItem, navCollapseThreshold()+1)
	if !NavSectionNeedsCollapse(long) {
		t.Fatal("expected collapse above threshold")
	}
}

func TestNavSectionExpanded(t *testing.T) {
	items := make([]NavItem, 12)
	for i := range items {
		items[i] = NavItem{Path: fmt.Sprintf("/p/%d/", i)}
	}
	if !NavSectionExpanded(items[5].Path, items) {
		t.Fatal("expected expanded when active is in fade row")
	}
	if !NavSectionExpanded(items[10].Path, items) {
		t.Fatal("expected expanded when active is hidden")
	}
	if NavSectionExpanded(items[2].Path, items) {
		t.Fatal("expected collapsed when active in first five")
	}
}

func TestNavSectionCollapseID(t *testing.T) {
	if got := NavSectionCollapseID("Components", false); got != "nav-collapse-components" {
		t.Fatalf("got %q want nav-collapse-components", got)
	}
	if got := NavSectionCollapseID("Components", true); got != "nav-collapse-components-mobile" {
		t.Fatalf("got %q want nav-collapse-components-mobile", got)
	}
	if got := NavSectionCollapseID("", false); got != "nav-collapse-section" {
		t.Fatalf("empty label: got %q", got)
	}
}
