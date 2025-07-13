package StencilPage

import (
	"fmt"
	"strings"
)

// Ul génère une liste non ordonnée <ul>
func Ul(items []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var listItems []string
	for _, item := range items {
		listItems = append(listItems, fmt.Sprintf("<li>%s</li>", item))
	}

	content := strings.Join(listItems, "")
	return fmt.Sprintf(`<ul%s>%s</ul>`, classAttr, content)
}

// Ol génère une liste ordonnée <ol>
func Ol(items []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var listItems []string
	for _, item := range items {
		listItems = append(listItems, fmt.Sprintf("<li>%s</li>", item))
	}

	content := strings.Join(listItems, "")
	return fmt.Sprintf(`<ol%s>%s</ol>`, classAttr, content)
}

// Li génère un élément de liste <li>
func Li(content string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<li%s>%s</li>`, classAttr, content)
}

// Dl génère une liste de définitions <dl>
func Dl(items map[string]string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var content []string
	for term, definition := range items {
		content = append(content, fmt.Sprintf("<dt>%s</dt><dd>%s</dd>", term, definition))
	}

	contentStr := strings.Join(content, "")
	return fmt.Sprintf(`<dl%s>%s</dl>`, classAttr, contentStr)
}
