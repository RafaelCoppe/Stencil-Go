package StencilPage

import (
	"fmt"
	"strings"
)

// Table génère une table HTML complète
func Table(headers []string, rows [][]string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	// Générer les en-têtes
	var headerCells []string
	for _, header := range headers {
		headerCells = append(headerCells, fmt.Sprintf("<th>%s</th>", header))
	}
	headerRow := fmt.Sprintf("<tr>%s</tr>", strings.Join(headerCells, ""))
	thead := fmt.Sprintf("<thead>%s</thead>", headerRow)

	// Générer les lignes
	var bodyRows []string
	for _, row := range rows {
		var cells []string
		for _, cell := range row {
			cells = append(cells, fmt.Sprintf("<td>%s</td>", cell))
		}
		bodyRows = append(bodyRows, fmt.Sprintf("<tr>%s</tr>", strings.Join(cells, "")))
	}
	tbody := fmt.Sprintf("<tbody>%s</tbody>", strings.Join(bodyRows, ""))

	return fmt.Sprintf(`<table%s>%s%s</table>`, classAttr, thead, tbody)
}

// TableSimple génère une table simple sans en-têtes
func TableSimple(rows [][]string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var bodyRows []string
	for _, row := range rows {
		var cells []string
		for _, cell := range row {
			cells = append(cells, fmt.Sprintf("<td>%s</td>", cell))
		}
		bodyRows = append(bodyRows, fmt.Sprintf("<tr>%s</tr>", strings.Join(cells, "")))
	}

	return fmt.Sprintf(`<table%s><tbody>%s</tbody></table>`, classAttr, strings.Join(bodyRows, ""))
}

// Tr génère une ligne de table <tr>
func Tr(cells []string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var cellElements []string
	for _, cell := range cells {
		cellElements = append(cellElements, fmt.Sprintf("<td>%s</td>", cell))
	}

	return fmt.Sprintf(`<tr%s>%s</tr>`, classAttr, strings.Join(cellElements, ""))
}

// Td génère une cellule de table <td>
func Td(content string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<td%s>%s</td>`, classAttr, content)
}

// Th génère une cellule d'en-tête <th>
func Th(content string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<th%s>%s</th>`, classAttr, content)
}
