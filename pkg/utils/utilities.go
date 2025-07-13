package StencilUtils

import (
	"fmt"
	"strings"
)

// Br génère un saut de ligne <br />
func Br() string {
	return "<br />"
}

// Nbsp génère un espace insécable
func Nbsp() string {
	return "&nbsp;"
}

// Comment génère un commentaire HTML
func Comment(text string) string {
	return fmt.Sprintf("<!-- %s -->", text)
}

// Attribute génère un attribut HTML formaté
func Attribute(name, value string) string {
	if value == "" {
		return name
	}
	return fmt.Sprintf(`%s="%s"`, name, value)
}

// Attributes génère plusieurs attributs HTML
func Attributes(attrs map[string]string) string {
	var result []string
	for name, value := range attrs {
		result = append(result, Attribute(name, value))
	}
	return strings.Join(result, " ")
}

// EscapeHTML échappe les caractères HTML spéciaux
func EscapeHTML(text string) string {
	text = strings.ReplaceAll(text, "&", "&amp;")
	text = strings.ReplaceAll(text, "<", "&lt;")
	text = strings.ReplaceAll(text, ">", "&gt;")
	text = strings.ReplaceAll(text, "\"", "&quot;")
	text = strings.ReplaceAll(text, "'", "&#39;")
	return text
}

// ClassBuilder construit une chaîne de classes à partir d'un map conditionnel
func ClassBuilder(classes map[string]bool) string {
	var result []string
	for class, condition := range classes {
		if condition {
			result = append(result, class)
		}
	}
	return strings.Join(result, " ")
}

// Wrap génère un élément HTML avec du contenu
func Wrap(tag, content string, attributes map[string]string, classes ...string) string {
	var attrs []string

	// Ajouter les classes
	if len(classes) > 0 {
		attrs = append(attrs, fmt.Sprintf(`class="%s"`, strings.Join(classes, " ")))
	}

	// Ajouter les autres attributs
	for name, value := range attributes {
		if name != "class" { // Éviter la duplication de class
			attrs = append(attrs, Attribute(name, value))
		}
	}

	attrStr := ""
	if len(attrs) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	return fmt.Sprintf("<%s%s>%s</%s>", tag, attrStr, content, tag)
}

// SelfClosingTag génère un élément HTML auto-fermant
func SelfClosingTag(tag string, attributes map[string]string, classes ...string) string {
	var attrs []string

	// Ajouter les classes
	if len(classes) > 0 {
		attrs = append(attrs, fmt.Sprintf(`class="%s"`, strings.Join(classes, " ")))
	}

	// Ajouter les autres attributs
	for name, value := range attributes {
		if name != "class" {
			attrs = append(attrs, Attribute(name, value))
		}
	}

	attrStr := ""
	if len(attrs) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	return fmt.Sprintf("<%s%s />", tag, attrStr)
}
