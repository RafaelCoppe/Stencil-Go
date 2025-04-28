package StencilPage

import (
	"testing"
	"strings"
	"github.com/RafaelCoppe/Stencil-Go/pkg/page"
)
func TestDiv(t *testing.T) {
	children := []string{"<p>Test</p>"}
	extraClasses := []string{"bg-red-500", "p-4"}

	expected := `<div class="bg-red-500 p-4"><p>Test</p></div>`
	result := StencilPage.Div(children, extraClasses)

	if result != expected {
		t.Errorf("Div() = %v, want %v", result, expected)
	}
}

func TestContainer(t *testing.T) {
	children := []string{"<p>Container content</p>"}
	extraClasses := []string{"custom-class"}

	result := StencilPage.Container(children, extraClasses)

	if !strings.Contains(result, "mx-auto") || !strings.Contains(result, "max-w-7xl") {
		t.Errorf("Container() missing base classes")
	}

	if !strings.Contains(result, "custom-class") {
		t.Errorf("Container() missing custom classes")
	}

	if !strings.Contains(result, "<p>Container content</p>") {
		t.Errorf("Container() missing children")
	}
}

func TestFlexBox(t *testing.T) {
	children := []string{"<span>Item1</span>", "<span>Item2</span>"}
	extraClasses := []string{"bg-gray-200"}
	direction := "flex-row"
	justify := "justify-between"
	align := "items-center"
	gap := "4"

	result := StencilPage.FlexBox(children, extraClasses, direction, justify, align, gap)

	if !strings.Contains(result, "flex") ||
		!strings.Contains(result, direction) ||
		!strings.Contains(result, justify) ||
		!strings.Contains(result, align) ||
		!strings.Contains(result, "gap-4") {
		t.Errorf("FlexBox() missing flex properties")
	}

	if !strings.Contains(result, "bg-gray-200") {
		t.Errorf("FlexBox() missing extra classes")
	}

	if !strings.Contains(result, "<span>Item1</span>") || !strings.Contains(result, "<span>Item2</span>") {
		t.Errorf("FlexBox() missing children")
	}
}

func TestPage(t *testing.T) {
	bodyElements := []string{"<div>Content</div>"}
	title := "My Page"

	result := StencilPage.Page(title, bodyElements)

	if !strings.Contains(result, "<!DOCTYPE html>") ||
		!strings.Contains(result, "<title>My Page</title>") ||
		!strings.Contains(result, "<div>Content</div>") {
		t.Errorf("Page() missing basic HTML structure or content")
	}
}
