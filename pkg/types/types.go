package StencilTypes

// HTMLElement represents a basic HTML element structure
type HTMLElement struct {
	Tag        string
	Content    string
	Attributes map[string]string
	Classes    []string
	Children   []HTMLElement
}

// FormElement represents form input types
type FormElement struct {
	Type        string
	Name        string
	Value       string
	Placeholder string
	Required    bool
	Disabled    bool
}

// LinkTarget represents target options for links
type LinkTarget string

const (
	TargetBlank  LinkTarget = "_blank"
	TargetSelf   LinkTarget = "_self"
	TargetParent LinkTarget = "_parent"
	TargetTop    LinkTarget = "_top"
)

// ButtonType represents button types
type ButtonType string

const (
	ButtonTypeButton ButtonType = "button"
	ButtonTypeSubmit ButtonType = "submit"
	ButtonTypeReset  ButtonType = "reset"
)

// InputType represents input field types
type InputType string

const (
	InputTypeText     InputType = "text"
	InputTypeEmail    InputType = "email"
	InputTypePassword InputType = "password"
	InputTypeNumber   InputType = "number"
	InputTypeDate     InputType = "date"
	InputTypeTime     InputType = "time"
	InputTypeFile     InputType = "file"
	InputTypeHidden   InputType = "hidden"
	InputTypeCheckbox InputType = "checkbox"
	InputTypeRadio    InputType = "radio"
)

// MediaControls represents media control options
type MediaControls struct {
	Controls bool
	Autoplay bool
	Loop     bool
	Muted    bool
}
