package StencilText

import StencilInsiders "github.com/RafaelCoppe/Stencil-Go/pkg/inside_functions"

func Paragraphe(text string, classes ...string) string {
	return StencilInsiders.Wrap("p", text, classes...)
}
