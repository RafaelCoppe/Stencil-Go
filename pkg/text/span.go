package StencilText

import StencilInsiders "github.com/RafaelCoppe/Stencil-Go/pkg/inside_functions"

func Span(text string, classes ...string) string {
	return StencilInsiders.Wrap("span", text, classes...)
}
