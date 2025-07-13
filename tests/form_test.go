package tests

import (
	"strings"
	"testing"

	StencilForm "github.com/RafaelCoppe/Stencil-Go/pkg/form"
)

func TestForm(t *testing.T) {
	form := StencilForm.Form("/submit", "POST", []string{"<input>test</input>"}, "form-class")

	if !strings.Contains(form, `<form action="/submit" method="POST"`) {
		t.Error("Form should contain action and method attributes")
	}

	if !strings.Contains(form, `class="form-class"`) {
		t.Error("Form should contain the specified class")
	}

	if !strings.Contains(form, "<input>test</input>") {
		t.Error("Form should contain the provided elements")
	}
}

func TestInputText(t *testing.T) {
	input := StencilForm.InputText("username", "Enter username", "john", "form-control")

	if !strings.Contains(input, `type="text"`) {
		t.Error("Input should be of type text")
	}

	if !strings.Contains(input, `name="username"`) {
		t.Error("Input should have the correct name")
	}

	if !strings.Contains(input, `placeholder="Enter username"`) {
		t.Error("Input should have the correct placeholder")
	}

	if !strings.Contains(input, `value="john"`) {
		t.Error("Input should have the correct value")
	}

	if !strings.Contains(input, `class="form-control"`) {
		t.Error("Input should have the correct class")
	}
}

func TestCheckbox(t *testing.T) {
	checkbox := StencilForm.Checkbox("agree", "yes", "I agree", true, "form-check")

	if !strings.Contains(checkbox, `type="checkbox"`) {
		t.Error("Checkbox should be of type checkbox")
	}

	if !strings.Contains(checkbox, `name="agree"`) {
		t.Error("Checkbox should have the correct name")
	}

	if !strings.Contains(checkbox, `value="yes"`) {
		t.Error("Checkbox should have the correct value")
	}

	if !strings.Contains(checkbox, "checked") {
		t.Error("Checkbox should be checked")
	}

	if !strings.Contains(checkbox, "I agree") {
		t.Error("Checkbox should have the correct label")
	}
}

func TestTextArea(t *testing.T) {
	textarea := StencilForm.TextArea("message", "Enter message", "Hello", 5, 30, "form-control")

	if !strings.Contains(textarea, `name="message"`) {
		t.Error("TextArea should have the correct name")
	}

	if !strings.Contains(textarea, `placeholder="Enter message"`) {
		t.Error("TextArea should have the correct placeholder")
	}

	if !strings.Contains(textarea, `rows="5"`) {
		t.Error("TextArea should have the correct rows")
	}

	if !strings.Contains(textarea, `cols="30"`) {
		t.Error("TextArea should have the correct cols")
	}

	if !strings.Contains(textarea, "Hello") {
		t.Error("TextArea should contain the correct content")
	}
}

func TestSelect(t *testing.T) {
	options := map[string]string{
		"fr": "Français",
		"en": "English",
		"es": "Español",
	}

	selectElement := StencilForm.Select("language", options, "fr", "form-select")

	if !strings.Contains(selectElement, `name="language"`) {
		t.Error("Select should have the correct name")
	}

	if !strings.Contains(selectElement, `class="form-select"`) {
		t.Error("Select should have the correct class")
	}

	if !strings.Contains(selectElement, `value="fr" selected`) {
		t.Error("Select should have the correct selected option")
	}

	if !strings.Contains(selectElement, "Français") {
		t.Error("Select should contain the option labels")
	}
}
