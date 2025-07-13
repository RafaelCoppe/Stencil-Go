package StencilForm

import (
	"fmt"
	"strings"
)

// RadioButton génère un bouton radio.
func RadioButton(name, value, labelText string, checked bool, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	checkedAttr := ""
	if checked {
		checkedAttr = " checked"
	}

	input := fmt.Sprintf(`<input type="radio" name="%s" value="%s"%s%s />`,
		name, value, classAttr, checkedAttr)

	return fmt.Sprintf(`<label>%s %s</label>`, input, labelText)
}

// RadioGroup génère un groupe de boutons radio.
func RadioGroup(name string, options []map[string]interface{}, classes ...string) string {
	var radioButtons []string

	for _, option := range options {
		value := option["value"].(string)
		label := option["label"].(string)
		checked := false
		if option["checked"] != nil {
			checked = option["checked"].(bool)
		}

		radioButtons = append(radioButtons, RadioButton(name, value, label, checked))
	}

	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<div%s>%s</div>`, classAttr, strings.Join(radioButtons, ""))
}

// FieldSet génère un fieldset avec une légende.
func FieldSet(legend string, content string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	legendElement := ""
	if legend != "" {
		legendElement = fmt.Sprintf(`<legend>%s</legend>`, legend)
	}

	return fmt.Sprintf(`<fieldset%s>%s%s</fieldset>`, classAttr, legendElement, content)
}

// FormGroup génère un groupe de formulaire avec label et input.
func FormGroup(labelText, inputHTML string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	label := ""
	if labelText != "" {
		label = fmt.Sprintf(`<label>%s</label>`, labelText)
	}

	return fmt.Sprintf(`<div%s>%s%s</div>`, classAttr, label, inputHTML)
}

// Range génère un input de type range (slider).
func Range(name string, min, max, value int, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<input type="range" name="%s" min="%d" max="%d" value="%d"%s />`,
		name, min, max, value, classAttr)
}

// Color génère un input de type color.
func Color(name, value string, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	valueAttr := ""
	if value != "" {
		valueAttr = fmt.Sprintf(` value="%s"`, value)
	}

	return fmt.Sprintf(`<input type="color" name="%s"%s%s />`,
		name, valueAttr, classAttr)
}

// Hidden génère un input de type hidden.
func Hidden(name, value string) string {
	return fmt.Sprintf(`<input type="hidden" name="%s" value="%s" />`, name, value)
}
