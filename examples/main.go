package main

import (
	"fmt"
	"os"

	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

func main() {
	// Contenu principal
	content := StencilUtils.Join(
		StencilText.Titre1("Bienvenue sur Stencil-Go", "text-center", "text-3xl", "mb-4"),
		StencilUtils.Hr("my-4"),
		StencilText.Paragraphe("Stencil-Go vous permet de générer facilement des éléments HTML avec Go, sans écrire de HTML à la main.", "text-lg"),
		StencilPage.Div(
			StencilText.Paragraphe("Ceci est une div imbriquée avec un paragraphe.", "text-sm", "text-gray-600"),
			"p-4", "bg-gray-100", "rounded",
		),
	)

	// Inclusion dans une container + page HTML complète
	fullContent := StencilPage.Container(content, "mx-auto", "mt-10")
	html := StencilPage.Page("Page de démonstration", []string{fullContent})

	// Écriture dans un fichier
	file, err := os.Create("output.html")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(html)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}

	fmt.Println("✅ Fichier output.html généré avec succès !")
}
