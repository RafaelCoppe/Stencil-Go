package tests

import (
	"strings"
	"testing"

	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
)

func TestPage(t *testing.T) {
	result := StencilPage.Page("Test Page", []string{"<h1>Title</h1>"})
	expected := "<html><head><title>Test Page</title></head><body><h1>Title</h1></body></html>"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestDiv(t *testing.T) {
	div := StencilPage.Div("Content inside div", "container", "my-4")

	if !strings.Contains(div, "<div") {
		t.Error("Div should generate a div tag")
	}

	if !strings.Contains(div, "</div>") {
		t.Error("Div should close properly")
	}

	if !strings.Contains(div, "Content inside div") {
		t.Error("Div should contain the content")
	}

	if !strings.Contains(div, "container my-4") {
		t.Error("Div should contain multiple classes")
	}
}

func TestContainer(t *testing.T) {
	container := StencilPage.Container("Page content", "mx-auto")

	if !strings.Contains(container, "<div") {
		t.Error("Container should generate a div tag")
	}

	if !strings.Contains(container, `class="container mx-auto"`) {
		t.Error("Container should include base container class and custom classes")
	}

	if !strings.Contains(container, "Page content") {
		t.Error("Container should contain the content")
	}
}

func TestUl(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}
	ul := StencilPage.Ul(items, "list-group")

	if !strings.Contains(ul, "<ul") {
		t.Error("Ul should generate a ul tag")
	}

	if !strings.Contains(ul, "</ul>") {
		t.Error("Ul should close properly")
	}

	if !strings.Contains(ul, `class="list-group"`) {
		t.Error("Ul should contain the class")
	}

	if !strings.Contains(ul, "<li>Item 1</li>") {
		t.Error("Ul should contain list items")
	}

	if !strings.Contains(ul, "<li>Item 2</li>") {
		t.Error("Ul should contain all list items")
	}
}

func TestOl(t *testing.T) {
	items := []string{"First", "Second", "Third"}
	ol := StencilPage.Ol(items, "numbered-list")

	if !strings.Contains(ol, "<ol") {
		t.Error("Ol should generate an ol tag")
	}

	if !strings.Contains(ol, "</ol>") {
		t.Error("Ol should close properly")
	}

	if !strings.Contains(ol, `class="numbered-list"`) {
		t.Error("Ol should contain the class")
	}

	if !strings.Contains(ol, "<li>First</li>") {
		t.Error("Ol should contain list items")
	}
}

func TestTable(t *testing.T) {
	headers := []string{"Name", "Age", "City"}
	rows := [][]string{
		{"John", "25", "Paris"},
		{"Jane", "30", "London"},
	}

	table := StencilPage.Table(headers, rows, "table", "table-striped")

	if !strings.Contains(table, "<table") {
		t.Error("Table should generate a table tag")
	}

	if !strings.Contains(table, "</table>") {
		t.Error("Table should close properly")
	}

	if !strings.Contains(table, `class="table table-striped"`) {
		t.Error("Table should contain the classes")
	}

	if !strings.Contains(table, "<thead>") {
		t.Error("Table should contain a thead")
	}

	if !strings.Contains(table, "<tbody>") {
		t.Error("Table should contain a tbody")
	}

	if !strings.Contains(table, "<th>Name</th>") {
		t.Error("Table should contain header cells")
	}

	if !strings.Contains(table, "<td>John</td>") {
		t.Error("Table should contain data cells")
	}
}
