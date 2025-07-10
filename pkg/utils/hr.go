package StencilUtils

import (
	"strings"
)

// Hr génère une balise <hr /> avec des classes éventuelles.
func Hr(classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = ` class="` + strings.Join(classes, " ") + `"`
	}
	return `<hr` + classAttr + ` />`
}
