package StencilInsiders

import (
	"fmt"
	"strings"
)

func Wrap(tag, text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}
	return fmt.Sprintf(`<%s%s>%s</%s>`, tag, classAttr, text, tag)
}