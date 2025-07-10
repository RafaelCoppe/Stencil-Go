package StencilForm

import (
	"fmt"
	"strings"
)

// Form génère un formulaire HTML complet.
func Form(action, method string, elements []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}
	body := strings.Join(elements, "\n")
	return fmt.Sprintf(`<form action="%s" method="%s"%s>%s</form>`, action, method, classAttr, body)
}