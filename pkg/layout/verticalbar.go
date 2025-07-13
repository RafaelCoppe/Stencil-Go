package StencilLayout

import (
	"fmt"
	"strings"
)

// VerticalBar génère une div avec display flex en mode column (vertical).
func VerticalBar(children []string, classes ...string) string {
	baseClasses := []string{"d-flex", "flex-column"}
	allClasses := append(baseClasses, classes...)

	classAttr := ""
	if len(allClasses) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(allClasses, " "))
	}

	content := strings.Join(children, "")
	return fmt.Sprintf(`<div%s>%s</div>`, classAttr, content)
}
