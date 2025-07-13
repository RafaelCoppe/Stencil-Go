package StencilInteractions

import (
	"fmt"
	"strings"
)

// Button génère un bouton HTML.
func Bouton(text string, event string, classes ...string) string {
	return fmt.Sprintf(`<button data-onclick="%s" class="%s">%s</button>`,
		event,
		strings.Join(classes, " "),
		text,
	)
}