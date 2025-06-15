package StencilText

import StencilInsiders "github.com/RafaelCoppe/Stencil-Go/pkg/inside_functions"

func Titre3(text string, classes ...string) string {
	return StencilInsiders.Wrap("h3", text, classes...)
}
