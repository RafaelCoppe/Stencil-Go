package StencilInteractions

import (
	"fmt"
	"github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// Button génère un bouton HTML.
func Bouton(label string, event string, classes []string) string {
	return fmt.Sprintf(`<button data-onclick="%s" class="%s">%s</button>`,
		event,
		StencilUtils.Join(classes...),
		label,
	)
}