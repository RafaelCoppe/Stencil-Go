package StencilText

import (
	StencilInsiders "github.com/RafaelCoppe/Stencil-Go/pkg/inside_functions"
)

func Titre2(text string, classes ...string) string {
	return StencilInsiders.Wrap("h2", text, classes...)
}
