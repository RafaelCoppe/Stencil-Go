package StencilPage

import (
	"fmt"
	"strings"
)

// Container génère une <div> avec une classe container.
func Container(content string, classes ...string) string {
	baseClass := "container"
	allClasses := append([]string{baseClass}, classes...)
	return fmt.Sprintf(`<div class="%s">%s</div>`, strings.Join(allClasses, " "), content)
}
