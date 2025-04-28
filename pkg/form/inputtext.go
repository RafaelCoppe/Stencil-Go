package form

import (
	"fmt"
	"strings"
)

// Input génère un champ input HTML.
func Input(inputType, name, placeholder string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}
	return fmt.Sprintf(`<input type="%s" name="%s" placeholder="%s"%s />`, inputType, name, placeholder, classAttr)
}