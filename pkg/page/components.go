package StencilPage

import (
	"fmt"
	"strings"
)

// Card génère une carte Bootstrap-style
func Card(title, content string, classes ...string) string {
	classAttr := ""
	baseClasses := []string{"card"}
	allClasses := append(baseClasses, classes...)

	if len(allClasses) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(allClasses, " "))
	}

	cardBody := ""
	if title != "" {
		cardBody = fmt.Sprintf(`<div class="card-body"><h5 class="card-title">%s</h5><div class="card-text">%s</div></div>`, title, content)
	} else {
		cardBody = fmt.Sprintf(`<div class="card-body">%s</div>`, content)
	}

	return fmt.Sprintf(`<div%s>%s</div>`, classAttr, cardBody)
}

// Alert génère une alerte
func Alert(message, alertType string, dismissible bool, classes ...string) string {
	baseClasses := []string{"alert", fmt.Sprintf("alert-%s", alertType)}
	if dismissible {
		baseClasses = append(baseClasses, "alert-dismissible")
	}
	allClasses := append(baseClasses, classes...)

	classAttr := fmt.Sprintf(` class="%s"`, strings.Join(allClasses, " "))

	dismissButton := ""
	if dismissible {
		dismissButton = `<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`
	}

	return fmt.Sprintf(`<div%s role="alert">%s%s</div>`, classAttr, message, dismissButton)
}

// Badge génère un badge
func Badge(text, badgeType string, classes ...string) string {
	baseClasses := []string{"badge"}
	if badgeType != "" {
		baseClasses = append(baseClasses, fmt.Sprintf("bg-%s", badgeType))
	}
	allClasses := append(baseClasses, classes...)

	classAttr := fmt.Sprintf(` class="%s"`, strings.Join(allClasses, " "))

	return fmt.Sprintf(`<span%s>%s</span>`, classAttr, text)
}

// Modal génère un modal Bootstrap
func Modal(id, title, body, footer string, classes ...string) string {
	baseClasses := []string{"modal", "fade"}
	allClasses := append(baseClasses, classes...)

	classAttr := fmt.Sprintf(` class="%s"`, strings.Join(allClasses, " "))

	footerContent := ""
	if footer != "" {
		footerContent = fmt.Sprintf(`<div class="modal-footer">%s</div>`, footer)
	}

	return fmt.Sprintf(`<div%s id="%s" tabindex="-1" aria-labelledby="%sLabel" aria-hidden="true">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h1 class="modal-title fs-5" id="%sLabel">%s</h1>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">%s</div>
				%s
			</div>
		</div>
	</div>`, classAttr, id, id, id, title, body, footerContent)
}

// Accordion génère un accordéon
func Accordion(id string, items []map[string]string, classes ...string) string {
	baseClasses := []string{"accordion"}
	allClasses := append(baseClasses, classes...)

	classAttr := fmt.Sprintf(` class="%s"`, strings.Join(allClasses, " "))

	var accordionItems []string
	for i, item := range items {
		itemId := fmt.Sprintf("%s-item-%d", id, i)
		collapseId := fmt.Sprintf("%s-collapse-%d", id, i)

		expanded := ""
		show := ""
		if i == 0 { // Premier élément ouvert par défaut
			expanded = ` aria-expanded="true"`
			show = " show"
		} else {
			expanded = ` aria-expanded="false" class="collapsed"`
		}

		accordionItem := fmt.Sprintf(`
			<div class="accordion-item">
				<h2 class="accordion-header" id="%s">
					<button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#%s" aria-controls="%s"%s>
						%s
					</button>
				</h2>
				<div id="%s" class="accordion-collapse collapse%s" aria-labelledby="%s" data-bs-parent="#%s">
					<div class="accordion-body">%s</div>
				</div>
			</div>`, itemId, collapseId, collapseId, expanded, item["title"], collapseId, show, itemId, id, item["content"])

		accordionItems = append(accordionItems, accordionItem)
	}

	return fmt.Sprintf(`<div%s id="%s">%s</div>`, classAttr, id, strings.Join(accordionItems, ""))
}

// Progress génère une barre de progression
func Progress(value, max int, label string, classes ...string) string {
	baseClasses := []string{"progress"}
	allClasses := append(baseClasses, classes...)

	classAttr := fmt.Sprintf(` class="%s"`, strings.Join(allClasses, " "))

	percentage := 0
	if max > 0 {
		percentage = (value * 100) / max
	}

	progressLabel := ""
	if label != "" {
		progressLabel = label
	} else {
		progressLabel = fmt.Sprintf("%d%%", percentage)
	}

	return fmt.Sprintf(`<div%s role="progressbar" aria-label="%s" aria-valuenow="%d" aria-valuemin="0" aria-valuemax="%d">
		<div class="progress-bar" style="width: %d%%">%s</div>
	</div>`, classAttr, progressLabel, value, max, percentage, progressLabel)
}
