package StencilPage

import "strings"

// Div génère un élément <div> avec ses enfants et des classes CSS optionnelles.
func Div(children []string, extraClasses []string) string {
	classes := ""
	if len(extraClasses) > 0 {
		classes = ` class="` + strings.Join(extraClasses, " ") + `"`
	}
	return `<div` + classes + `>` + strings.Join(children, "") + `</div>`
}
