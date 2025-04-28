package StencilPage

import (
	"fmt"
	"strings"
)

// Page génère une page HTML complète avec un titre et un corps de page.
func Page(title string, bodyElements []string) string {
	bootstrap := `<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet">
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css">`
	body := fmt.Sprintf("<body>%s</body>", strings.Join(bodyElements, "\n"))
	return fmt.Sprintf("<!DOCTYPE html><html><head>%s<title>%s</title></head>%s</html>", bootstrap, title, body)
}
