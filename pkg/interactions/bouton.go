package StencilInteractions

import (
	"fmt"
	"strings"
)

// Bouton génère un bouton HTML.
func Bouton(text string, event string, classes ...string) string {
	return fmt.Sprintf(`<button data-onclick="%s" class="%s">%s</button>`,
		event,
		strings.Join(classes, " "),
		text,
	)
}

// BoutonSubmit génère un bouton de type submit.
func BoutonSubmit(text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<button type="submit"%s>%s</button>`, classAttr, text)
}

// BoutonReset génère un bouton de type reset.
func BoutonReset(text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<button type="reset"%s>%s</button>`, classAttr, text)
}

// BoutonDisabled génère un bouton désactivé.
func BoutonDisabled(text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<button disabled%s>%s</button>`, classAttr, text)
}

// BoutonWithId génère un bouton avec un ID spécifique.
func BoutonWithId(text, id, event string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<button id="%s" data-onclick="%s"%s>%s</button>`,
		id, event, classAttr, text)
}
