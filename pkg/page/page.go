package StencilPage

import (
	"fmt"
	"strings"
)

// Page génère une page HTML complète avec un titre et un corps de page.
func Page(title string, bodyElements []string) string {
	body := fmt.Sprintf("<body>%s</body>", strings.Join(bodyElements, "\n"))
	return fmt.Sprintf("<html><head><title>%s</title></head>%s</html>", title, body)
}
