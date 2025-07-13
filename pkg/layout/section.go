package StencilLayout

import (
	"fmt"
	"strings"
)

// Section génère une balise <section> HTML avec des classes optionnelles.
func Section(children []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	content := strings.Join(children, "")
	return fmt.Sprintf(`<section%s>%s</section>`, classAttr, content)
}

// Article génère une balise <article> HTML avec des classes optionnelles.
func Article(children []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	content := strings.Join(children, "")
	return fmt.Sprintf(`<article%s>%s</article>`, classAttr, content)
}

// Header génère une balise <header> HTML avec des classes optionnelles.
func Header(children []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	content := strings.Join(children, "")
	return fmt.Sprintf(`<header%s>%s</header>`, classAttr, content)
}

// Footer génère une balise <footer> HTML avec des classes optionnelles.
func Footer(children []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	content := strings.Join(children, "")
	return fmt.Sprintf(`<footer%s>%s</footer>`, classAttr, content)
}

// Nav génère une balise <nav> HTML avec des classes optionnelles.
func Nav(children []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	content := strings.Join(children, "")
	return fmt.Sprintf(`<nav%s>%s</nav>`, classAttr, content)
}

// Main génère une balise <main> HTML avec des classes optionnelles.
func Main(children []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	content := strings.Join(children, "")
	return fmt.Sprintf(`<main%s>%s</main>`, classAttr, content)
}
