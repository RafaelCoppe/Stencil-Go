package interactions

import (
	"fmt"
	"strings"
)

// Lien génère un lien HTML (<a>).
func Lien(href, text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<a href="%s"%s>%s</a>`, href, classAttr, text)
}
