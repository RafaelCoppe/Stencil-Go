package StencilPage

// Container génère un conteneur centré avec largeur maximale, basé sur Div.
func Container(children []string, extraClasses []string) string {
	// Classes de base pour un conteneur centré
	baseClasses := []string{"mx-auto", "max-w-7xl", "px-4", "sm:px-6", "lg:px-8"}
	// Ajout des classes supplémentaires
	allClasses := append(baseClasses, extraClasses...)
	return Div(children, allClasses)
}
	