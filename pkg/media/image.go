package StencilMedia

import (
	"fmt"
	"strings"
)

// Image génère une balise <img> HTML avec des classes optionnelles.
func Image(src, alt string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<img src="%s" alt="%s"%s />`, src, alt, classAttr)
}

// ImageWithTitle génère une image avec un titre (title attribute).
func ImageWithTitle(src, alt, title string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<img src="%s" alt="%s" title="%s"%s />`, src, alt, title, classAttr)
}

// ResponsiveImage génère une image responsive avec srcset.
func ResponsiveImage(src, srcset, alt string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	srcsetAttr := ""
	if srcset != "" {
		srcsetAttr = fmt.Sprintf(` srcset="%s"`, srcset)
	}

	return fmt.Sprintf(`<img src="%s"%s alt="%s"%s />`, src, srcsetAttr, alt, classAttr)
}

// Figure génère une balise <figure> avec une image et une légende.
func Figure(src, alt, caption string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	img := Image(src, alt)
	figcaption := ""
	if caption != "" {
		figcaption = fmt.Sprintf(`<figcaption>%s</figcaption>`, caption)
	}

	return fmt.Sprintf(`<figure%s>%s%s</figure>`, classAttr, img, figcaption)
}
