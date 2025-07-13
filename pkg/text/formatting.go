package StencilText

import StencilInsiders "github.com/RafaelCoppe/Stencil-Go/pkg/inside_functions"

// Code génère un élément <code> pour du code inline
func Code(text string, classes ...string) string {
	return StencilInsiders.Wrap("code", text, classes...)
}

// Pre génère un élément <pre> pour du code préformaté
func Pre(text string, classes ...string) string {
	return StencilInsiders.Wrap("pre", text, classes...)
}

// CodeBlock génère un bloc de code avec <pre><code>
func CodeBlock(text string, classes ...string) string {
	code := StencilInsiders.Wrap("code", text)
	return StencilInsiders.Wrap("pre", code, classes...)
}

// Strong génère un élément <strong> pour du texte en gras
func Strong(text string, classes ...string) string {
	return StencilInsiders.Wrap("strong", text, classes...)
}

// Em génère un élément <em> pour du texte en italique
func Em(text string, classes ...string) string {
	return StencilInsiders.Wrap("em", text, classes...)
}

// Small génère un élément <small> pour du petit texte
func Small(text string, classes ...string) string {
	return StencilInsiders.Wrap("small", text, classes...)
}

// Mark génère un élément <mark> pour du texte surligné
func Mark(text string, classes ...string) string {
	return StencilInsiders.Wrap("mark", text, classes...)
}

// Del génère un élément <del> pour du texte barré
func Del(text string, classes ...string) string {
	return StencilInsiders.Wrap("del", text, classes...)
}

// Ins génère un élément <ins> pour du texte souligné
func Ins(text string, classes ...string) string {
	return StencilInsiders.Wrap("ins", text, classes...)
}

// Blockquote génère un élément <blockquote> pour les citations
func Blockquote(text string, classes ...string) string {
	return StencilInsiders.Wrap("blockquote", text, classes...)
}
