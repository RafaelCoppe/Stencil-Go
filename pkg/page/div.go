package StencilPage

import (
	"fmt"
	"strings"
)

// Div génère une <div> simple avec des classes supplémentaires.
func Div(content string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}
	return fmt.Sprintf(`<div%s>%s</div>`, classAttr, content)
}
