package layout

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestSidebarNavCollapse_rendersTeaserIcons(t *testing.T) {
	t.Helper()
	items := make([]NavItem, navCollapseThreshold()+2)
	for i := range items {
		items[i] = NavItem{
			Label: fmt.Sprintf("Item%d", i),
			Path:  fmt.Sprintf("/p/%d/", i),
			Icon:  "layout-grid",
		}
	}
	group := NavSectionGroup{Label: "Primitives", Items: items}
	var buf bytes.Buffer
	if err := sidebarNavCollapse(SidebarProps{Active: ""}, group).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	if !strings.Contains(html, "nav-sidebar-teaser") {
		t.Fatal("expected teaser block")
	}
	teaserStart := strings.Index(html, "nav-sidebar-teaser")
	teaserEnd := strings.Index(html[teaserStart:], "nav-sidebar-overflow")
	if teaserEnd < 0 {
		t.Fatal("expected overflow block after teaser")
	}
	teaserHTML := html[teaserStart : teaserStart+teaserEnd]
	if !strings.Contains(teaserHTML, "latty-chevron-right") {
		t.Fatalf("expected teaser icons, got: %q", teaserHTML)
	}
}
