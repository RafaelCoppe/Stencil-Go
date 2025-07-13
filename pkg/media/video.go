package StencilMedia

import (
	"fmt"
	"strings"
)

// Video génère une balise <video> HTML avec des contrôles optionnels.
func Video(src string, controls bool, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	controlsAttr := ""
	if controls {
		controlsAttr = " controls"
	}

	return fmt.Sprintf(`<video src="%s"%s%s></video>`, src, controlsAttr, classAttr)
}

// VideoWithOptions génère une vidéo avec des options avancées.
func VideoWithOptions(src string, controls, autoplay, loop, muted bool, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	var attrs []string
	if controls {
		attrs = append(attrs, "controls")
	}
	if autoplay {
		attrs = append(attrs, "autoplay")
	}
	if loop {
		attrs = append(attrs, "loop")
	}
	if muted {
		attrs = append(attrs, "muted")
	}

	attrStr := ""
	if len(attrs) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	return fmt.Sprintf(`<video src="%s"%s%s></video>`, src, attrStr, classAttr)
}

// VideoWithSources génère une vidéo avec plusieurs sources.
func VideoWithSources(sources map[string]string, controls bool, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	controlsAttr := ""
	if controls {
		controlsAttr = " controls"
	}

	var sourceTags []string
	for format, src := range sources {
		sourceTags = append(sourceTags, fmt.Sprintf(`<source src="%s" type="video/%s">`, src, format))
	}

	sourceContent := strings.Join(sourceTags, "")

	return fmt.Sprintf(`<video%s%s>%s<p>Your browser doesn't support HTML5 video.</p></video>`,
		controlsAttr, classAttr, sourceContent)
}

// Audio génère une balise <audio> HTML.
func Audio(src string, controls bool, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	controlsAttr := ""
	if controls {
		controlsAttr = " controls"
	}

	return fmt.Sprintf(`<audio src="%s"%s%s></audio>`, src, controlsAttr, classAttr)
}

// Iframe génère une balise <iframe> HTML.
func Iframe(src string, width, height int, classes ...string) string {
	classAttr := ""
	if len(classes) > 0 {
		classAttr = fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<iframe src="%s" width="%d" height="%d"%s></iframe>`,
		src, width, height, classAttr)
}
