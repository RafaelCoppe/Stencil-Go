package StencilPage

// FlexBox génère une div flex personnalisable.
func FlexBox(children []string, extraClasses []string, direction string, justify string, align string, gap string) string {
	flexClasses := []string{"flex"}

	if direction != "" {
		flexClasses = append(flexClasses, direction)
	}
	if justify != "" {
		flexClasses = append(flexClasses, justify)
	}
	if align != "" {
		flexClasses = append(flexClasses, align)
	}
	if gap != "" {
		flexClasses = append(flexClasses, "gap-"+gap)
	}

	// Ajout des classes personnalisées
	flexClasses = append(flexClasses, extraClasses...)

	return Div(children, flexClasses)
}
