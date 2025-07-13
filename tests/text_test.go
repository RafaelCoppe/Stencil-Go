package tests

import (
	"strings"
	"testing"

	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
)

func TestTitre1(t *testing.T) {
	h1 := StencilText.Titre1("Hello World", "title-class")

	if !strings.Contains(h1, "<h1") {
		t.Error("Titre1 should generate an h1 tag")
	}

	if !strings.Contains(h1, "</h1>") {
		t.Error("Titre1 should close properly")
	}

	if !strings.Contains(h1, "Hello World") {
		t.Error("Titre1 should contain the text")
	}

	if !strings.Contains(h1, `class="title-class"`) {
		t.Error("Titre1 should contain the class")
	}
}

func TestTitre2(t *testing.T) {
	h2 := StencilText.Titre2("Subtitle", "subtitle-class", "mb-4")

	if !strings.Contains(h2, "<h2") {
		t.Error("Titre2 should generate an h2 tag")
	}

	if !strings.Contains(h2, "</h2>") {
		t.Error("Titre2 should close properly")
	}

	if !strings.Contains(h2, "Subtitle") {
		t.Error("Titre2 should contain the text")
	}

	if !strings.Contains(h2, "subtitle-class mb-4") {
		t.Error("Titre2 should contain multiple classes")
	}
}

func TestParagraphe(t *testing.T) {
	p := StencilText.Paragraphe("This is a paragraph", "text-muted")

	if !strings.Contains(p, "<p") {
		t.Error("Paragraphe should generate a p tag")
	}

	if !strings.Contains(p, "</p>") {
		t.Error("Paragraphe should close properly")
	}

	if !strings.Contains(p, "This is a paragraph") {
		t.Error("Paragraphe should contain the text")
	}

	if !strings.Contains(p, `class="text-muted"`) {
		t.Error("Paragraphe should contain the class")
	}
}

func TestSpan(t *testing.T) {
	span := StencilText.Span("Inline text", "badge", "bg-primary")

	if !strings.Contains(span, "<span") {
		t.Error("Span should generate a span tag")
	}

	if !strings.Contains(span, "</span>") {
		t.Error("Span should close properly")
	}

	if !strings.Contains(span, "Inline text") {
		t.Error("Span should contain the text")
	}

	if !strings.Contains(span, "badge bg-primary") {
		t.Error("Span should contain multiple classes")
	}
}

func TestCode(t *testing.T) {
	code := StencilText.Code("console.log('hello')", "language-js")

	if !strings.Contains(code, "<code") {
		t.Error("Code should generate a code tag")
	}

	if !strings.Contains(code, "</code>") {
		t.Error("Code should close properly")
	}

	if !strings.Contains(code, "console.log('hello')") {
		t.Error("Code should contain the code text")
	}

	if !strings.Contains(code, `class="language-js"`) {
		t.Error("Code should contain the class")
	}
}

func TestStrong(t *testing.T) {
	strong := StencilText.Strong("Important text", "fw-bold")

	if !strings.Contains(strong, "<strong") {
		t.Error("Strong should generate a strong tag")
	}

	if !strings.Contains(strong, "</strong>") {
		t.Error("Strong should close properly")
	}

	if !strings.Contains(strong, "Important text") {
		t.Error("Strong should contain the text")
	}
}
