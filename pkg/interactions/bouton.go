package StencilInteractions

import (
	"fmt"
	"strings"
)

// Button génère un bouton HTML.
func Button(buttonType, text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}
	return fmt.Sprintf(`<button type="%s"%s>%s</button>`, buttonType, classAttr, text)
}