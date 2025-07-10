package StencilForm

import (
	"fmt"
	"strings"
)

// Checkbox génère une case à cocher (checkbox) HTML.
func Checkbox(name, value, labelText string, checked bool, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	checkedAttr := ""
	if checked {
		checkedAttr = " checked"
	}

	input := fmt.Sprintf(`<input type="checkbox" name="%s" value="%s"%s%s />`, name, value, classAttr, checkedAttr)

	return fmt.Sprintf(`<label>%s %s</label>`, input, labelText)
}
