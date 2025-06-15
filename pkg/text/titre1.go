package StencilText

import StencilInsiders "github.com/RafaelCoppe/Stencil-Go/pkg/inside_functions"

func Titre1(text string, classes ...string) string {
	return StencilInsiders.Wrap("h1", text, classes...)
}
