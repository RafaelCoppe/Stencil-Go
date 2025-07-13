package StencilForm

import (
	"fmt"
	"strings"
)

// Input génère un champ input HTML.
func Input(inputType, name, placeholder string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}
	return fmt.Sprintf(`<input type="%s" name="%s" placeholder="%s"%s />`, inputType, name, placeholder, classAttr)
}

// InputText génère un champ input de type text avec une valeur.
func InputText(name, placeholder, value string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	valueAttr := ""
	if value != "" {
		valueAttr = fmt.Sprintf(` value="%s"`, value)
	}

	return fmt.Sprintf(`<input type="text" name="%s" placeholder="%s"%s%s />`,
		name, placeholder, valueAttr, classAttr)
}

// InputEmail génère un champ input de type email.
func InputEmail(name, placeholder, value string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	valueAttr := ""
	if value != "" {
		valueAttr = fmt.Sprintf(` value="%s"`, value)
	}

	return fmt.Sprintf(`<input type="email" name="%s" placeholder="%s"%s%s />`,
		name, placeholder, valueAttr, classAttr)
}

// InputPassword génère un champ input de type password.
func InputPassword(name, placeholder string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<input type="password" name="%s" placeholder="%s"%s />`,
		name, placeholder, classAttr)
}

// InputNumber génère un champ input de type number.
func InputNumber(name, placeholder string, value, min, max int, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var attrs []string
	if min != 0 {
		attrs = append(attrs, fmt.Sprintf(`min="%d"`, min))
	}
	if max != 0 {
		attrs = append(attrs, fmt.Sprintf(`max="%d"`, max))
	}

	attrStr := ""
	if len(attrs) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	return fmt.Sprintf(`<input type="number" name="%s" placeholder="%s" value="%d"%s%s />`,
		name, placeholder, value, attrStr, classAttr)
}

// InputDate génère un champ input de type date.
func InputDate(name string, value string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	valueAttr := ""
	if value != "" {
		valueAttr = fmt.Sprintf(` value="%s"`, value)
	}

	return fmt.Sprintf(`<input type="date" name="%s"%s%s />`,
		name, valueAttr, classAttr)
}

// InputFile génère un champ input de type file.
func InputFile(name string, accept string, multiple bool, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var attrs []string
	if accept != "" {
		attrs = append(attrs, fmt.Sprintf(`accept="%s"`, accept))
	}
	if multiple {
		attrs = append(attrs, "multiple")
	}

	attrStr := ""
	if len(attrs) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	return fmt.Sprintf(`<input type="file" name="%s"%s%s />`,
		name, attrStr, classAttr)
}

// TextArea génère un champ textarea HTML.
func TextArea(name, placeholder, value string, rows, cols int, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var attrs []string
	if rows > 0 {
		attrs = append(attrs, fmt.Sprintf(`rows="%d"`, rows))
	}
	if cols > 0 {
		attrs = append(attrs, fmt.Sprintf(`cols="%d"`, cols))
	}

	attrStr := ""
	if len(attrs) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	placeholderAttr := ""
	if placeholder != "" {
		placeholderAttr = fmt.Sprintf(` placeholder="%s"`, placeholder)
	}

	return fmt.Sprintf(`<textarea name="%s"%s%s%s>%s</textarea>`,
		name, placeholderAttr, attrStr, classAttr, value)
}

// Select génère un élément select avec des options.
func Select(name string, options map[string]string, selected string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var optionTags []string
	for value, label := range options {
		selectedAttr := ""
		if value == selected {
			selectedAttr = " selected"
		}
		optionTags = append(optionTags, fmt.Sprintf(`<option value="%s"%s>%s</option>`,
			value, selectedAttr, label))
	}

	optionsContent := strings.Join(optionTags, "")

	return fmt.Sprintf(`<select name="%s"%s>%s</select>`,
		name, classAttr, optionsContent)
}
