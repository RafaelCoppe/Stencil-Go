package tests

import (
	"strings"
	"testing"

	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
)

func TestBouton(t *testing.T) {
	button := StencilInteractions.Bouton("Click me", "handleClick", "btn", "btn-primary")

	if !strings.Contains(button, "<button") {
		t.Error("Bouton should generate a button tag")
	}

	if !strings.Contains(button, "</button>") {
		t.Error("Bouton should close properly")
	}

	if !strings.Contains(button, "Click me") {
		t.Error("Bouton should contain the text")
	}

	if !strings.Contains(button, `data-onclick="handleClick"`) {
		t.Error("Bouton should contain the event attribute")
	}

	if !strings.Contains(button, "btn btn-primary") {
		t.Error("Bouton should contain the classes")
	}
}

func TestBoutonSubmit(t *testing.T) {
	button := StencilInteractions.BoutonSubmit("Submit", "btn", "btn-success")

	if !strings.Contains(button, `type="submit"`) {
		t.Error("BoutonSubmit should be of type submit")
	}

	if !strings.Contains(button, "Submit") {
		t.Error("BoutonSubmit should contain the text")
	}

	if !strings.Contains(button, "btn btn-success") {
		t.Error("BoutonSubmit should contain the classes")
	}
}

func TestLien(t *testing.T) {
	link := StencilInteractions.Lien("/about", "About Us", "nav-link", "active")

	if !strings.Contains(link, "<a") {
		t.Error("Lien should generate an a tag")
	}

	if !strings.Contains(link, "</a>") {
		t.Error("Lien should close properly")
	}

	if !strings.Contains(link, `href="/about"`) {
		t.Error("Lien should contain the href attribute")
	}

	if !strings.Contains(link, "About Us") {
		t.Error("Lien should contain the text")
	}

	if !strings.Contains(link, "nav-link active") {
		t.Error("Lien should contain the classes")
	}
}

func TestExternalLink(t *testing.T) {
	link := StencilInteractions.ExternalLink("https://example.com", "External Site", "external-link")

	if !strings.Contains(link, `href="https://example.com"`) {
		t.Error("ExternalLink should contain the href attribute")
	}

	if !strings.Contains(link, `target="_blank"`) {
		t.Error("ExternalLink should open in new tab")
	}

	if !strings.Contains(link, `rel="noopener noreferrer"`) {
		t.Error("ExternalLink should have security attributes")
	}

	if !strings.Contains(link, "External Site") {
		t.Error("ExternalLink should contain the text")
	}
}

func TestMailtoLink(t *testing.T) {
	link := StencilInteractions.MailtoLink("test@example.com", "Contact Us", "mail-link")

	if !strings.Contains(link, `href="mailto:test@example.com"`) {
		t.Error("MailtoLink should contain mailto href")
	}

	if !strings.Contains(link, "Contact Us") {
		t.Error("MailtoLink should contain the text")
	}

	if !strings.Contains(link, `class="mail-link"`) {
		t.Error("MailtoLink should contain the class")
	}
}

func TestBreadcrumb(t *testing.T) {
	items := []map[string]string{
		{"href": "/", "text": "Home"},
		{"href": "/products", "text": "Products"},
		{"href": "", "text": "Current Page"},
	}

	breadcrumb := StencilInteractions.Breadcrumb(items, "breadcrumb")

	if !strings.Contains(breadcrumb, "<nav") {
		t.Error("Breadcrumb should generate a nav tag")
	}

	if !strings.Contains(breadcrumb, `class="breadcrumb"`) {
		t.Error("Breadcrumb should contain the class")
	}

	if !strings.Contains(breadcrumb, `<a href="/">Home</a>`) {
		t.Error("Breadcrumb should contain linked items")
	}

	if !strings.Contains(breadcrumb, "Current Page") {
		t.Error("Breadcrumb should contain the current page as text")
	}

	if !strings.Contains(breadcrumb, " / ") {
		t.Error("Breadcrumb should contain separators")
	}
}
