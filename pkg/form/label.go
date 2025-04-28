package form

import (
	"fmt"
	"strings"
)

// Label génère un label HTML associé à un champ.
func Label(forInput, text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}
	return fmt.Sprintf(`<label for="%s"%s>%s</label>`, forInput, classAttr, text)
}