package StencilUtils

import (
	"strings"
)

// Join concatène plusieurs morceaux de HTML.
func Join(parts ...string) string {
	return strings.Join(parts, "")
}
