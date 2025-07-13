package tests

import (
	"strings"
	"testing"

	StencilLayout "github.com/RafaelCoppe/Stencil-Go/pkg/layout"
)

func TestHorizontalBar(t *testing.T) {
	children := []string{"<div>Item 1</div>", "<div>Item 2</div>"}
	hBar := StencilLayout.HorizontalBar(children, "justify-content-between")

	if !strings.Contains(hBar, "d-flex") {
		t.Error("HorizontalBar should contain d-flex class")
	}

	if !strings.Contains(hBar, "flex-row") {
		t.Error("HorizontalBar should contain flex-row class")
	}

	if !strings.Contains(hBar, "justify-content-between") {
		t.Error("HorizontalBar should contain custom classes")
	}

	if !strings.Contains(hBar, "<div>Item 1</div>") {
		t.Error("HorizontalBar should contain all children")
	}

	if !strings.Contains(hBar, "<div>Item 2</div>") {
		t.Error("HorizontalBar should contain all children")
	}
}

func TestVerticalBar(t *testing.T) {
	children := []string{"<div>Item 1</div>", "<div>Item 2</div>"}
	vBar := StencilLayout.VerticalBar(children, "align-items-center")

	if !strings.Contains(vBar, "d-flex") {
		t.Error("VerticalBar should contain d-flex class")
	}

	if !strings.Contains(vBar, "flex-column") {
		t.Error("VerticalBar should contain flex-column class")
	}

	if !strings.Contains(vBar, "align-items-center") {
		t.Error("VerticalBar should contain custom classes")
	}

	if !strings.Contains(vBar, "<div>Item 1</div>") {
		t.Error("VerticalBar should contain all children")
	}
}

func TestSection(t *testing.T) {
	children := []string{"<h1>Title</h1>", "<p>Content</p>"}
	section := StencilLayout.Section(children, "my-section")

	if !strings.Contains(section, "<section") {
		t.Error("Section should generate a section tag")
	}

	if !strings.Contains(section, "</section>") {
		t.Error("Section should close properly")
	}

	if !strings.Contains(section, `class="my-section"`) {
		t.Error("Section should contain custom classes")
	}

	if !strings.Contains(section, "<h1>Title</h1>") {
		t.Error("Section should contain all children")
	}
}

func TestHeader(t *testing.T) {
	children := []string{"<nav>Navigation</nav>"}
	header := StencilLayout.Header(children, "main-header")

	if !strings.Contains(header, "<header") {
		t.Error("Header should generate a header tag")
	}

	if !strings.Contains(header, "</header>") {
		t.Error("Header should close properly")
	}

	if !strings.Contains(header, `class="main-header"`) {
		t.Error("Header should contain custom classes")
	}

	if !strings.Contains(header, "<nav>Navigation</nav>") {
		t.Error("Header should contain all children")
	}
}

func TestFooter(t *testing.T) {
	children := []string{"<p>Copyright</p>"}
	footer := StencilLayout.Footer(children, "main-footer")

	if !strings.Contains(footer, "<footer") {
		t.Error("Footer should generate a footer tag")
	}

	if !strings.Contains(footer, "</footer>") {
		t.Error("Footer should close properly")
	}

	if !strings.Contains(footer, `class="main-footer"`) {
		t.Error("Footer should contain custom classes")
	}

	if !strings.Contains(footer, "<p>Copyright</p>") {
		t.Error("Footer should contain all children")
	}
}
