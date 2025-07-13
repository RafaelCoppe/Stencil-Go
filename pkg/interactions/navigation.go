package StencilInteractions

import (
	"fmt"
	"strings"
)

// NavLink génère un lien de navigation avec état actif.
func NavLink(href, text string, active bool, classes ...string) string {
	baseClasses := []string{}
	if active {
		baseClasses = append(baseClasses, "active")
	}
	allClasses := append(baseClasses, classes...)

	classAttr := ""
	if len(allClasses) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(allClasses, " "))
	}

	return fmt.Sprintf(`<a href="%s"%s>%s</a>`, href, classAttr, text)
}

// ExternalLink génère un lien externe qui s'ouvre dans un nouvel onglet.
func ExternalLink(href, text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<a href="%s" target="_blank" rel="noopener noreferrer"%s>%s</a>`,
		href, classAttr, text)
}

// MailtoLink génère un lien mailto.
func MailtoLink(email, text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<a href="mailto:%s"%s>%s</a>`, email, classAttr, text)
}

// TelLink génère un lien téléphone.
func TelLink(phone, text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<a href="tel:%s"%s>%s</a>`, phone, classAttr, text)
}

// DownloadLink génère un lien de téléchargement.
func DownloadLink(href, filename, text string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	downloadAttr := ""
	if filename != "" {
		downloadAttr = fmt.Sprintf(` download="%s"`, filename)
	} else {
		downloadAttr = " download"
	}

	return fmt.Sprintf(`<a href="%s"%s%s>%s</a>`, href, downloadAttr, classAttr, text)
}

// Breadcrumb génère un fil d'Ariane.
func Breadcrumb(items []map[string]string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var breadcrumbItems []string
	for i, item := range items {
		href := item["href"]
		text := item["text"]

		if href != "" && i < len(items)-1 {
			// Lien normal pour tous les éléments sauf le dernier
			breadcrumbItems = append(breadcrumbItems,
				fmt.Sprintf(`<a href="%s">%s</a>`, href, text))
		} else {
			// Texte simple pour le dernier élément
			breadcrumbItems = append(breadcrumbItems, text)
		}
	}

	return fmt.Sprintf(`<nav%s>%s</nav>`, classAttr, strings.Join(breadcrumbItems, " / "))
}

// Pagination génère des liens de pagination.
func Pagination(currentPage, totalPages int, baseURL string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var links []string

	// Lien précédent
	if currentPage > 1 {
		links = append(links,
			fmt.Sprintf(`<a href="%s?page=%d">Previous</a>`, baseURL, currentPage-1))
	}

	// Liens des pages
	for i := 1; i <= totalPages; i++ {
		if i == currentPage {
			links = append(links, fmt.Sprintf(`<span class="current">%d</span>`, i))
		} else {
			links = append(links,
				fmt.Sprintf(`<a href="%s?page=%d">%d</a>`, baseURL, i, i))
		}
	}

	// Lien suivant
	if currentPage < totalPages {
		links = append(links,
			fmt.Sprintf(`<a href="%s?page=%d">Next</a>`, baseURL, currentPage+1))
	}

	return fmt.Sprintf(`<nav%s>%s</nav>`, classAttr, strings.Join(links, " "))
}
