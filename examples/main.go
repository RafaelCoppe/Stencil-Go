package main

import (
	"fmt"
	"os"
	"github.com/RafaelCoppe/Stencil-Go/pkg/page"
)

func main() {
	html := StencilPage.Page("Test Page", []string{"<h1>Title</h1>"})

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

	fmt.Println("Fichier output.html généré avec succès !")
}
