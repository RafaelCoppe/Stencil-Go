package StencilLayout

import (
	"fmt"
	"strings"
)

// HorizontalBar génère une div avec display flex en mode row (horizontal).
func HorizontalBar(children []string, classes ...string) string {
	baseClasses := []string{"d-flex", "flex-row"}
	allClasses := append(baseClasses, classes...)

	classAttr := ""
	if len(allClasses) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(allClasses, " "))
	}

	content := strings.Join(children, "")
	return fmt.Sprintf(`<div%s>%s</div>`, classAttr, content)
}
