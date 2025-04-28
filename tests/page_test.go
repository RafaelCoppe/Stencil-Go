package StencilPage

import (
	"testing"

	"github.com/RafaelCoppe/Stencil-Go/pkg/page"
)

func TestPage(t *testing.T) {
	result := StencilPage.Page("Test Page", []string{"<h1>Title</h1>"})
	expected := "<html><head><title>Test Page</title></head><body><h1>Title</h1></body></html>"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
