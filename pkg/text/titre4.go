package StencilText

import StencilInsiders "github.com/RafaelCoppe/Stencil-Go/pkg/inside_functions"

// Titre4 génère un titre H4
func Titre4(text string, classes ...string) string {
	return StencilInsiders.Wrap("h4", text, classes...)
}

// Titre5 génère un titre H5
func Titre5(text string, classes ...string) string {
	return StencilInsiders.Wrap("h5", text, classes...)
}

// Titre6 génère un titre H6
func Titre6(text string, classes ...string) string {
	return StencilInsiders.Wrap("h6", text, classes...)
}
