# 🖌️ Stencil - Générateur HTML déclaratif en Go

**Stencil** est une bibliothèque open-source écrite en Go qui permet de générer facilement du HTML sans avoir besoin de connaissances en front-end. Elle adopte une approche déclarative inspirée de SwiftUI : les éléments HTML sont créés via des fonctions imbriquées, avec gestion automatique des classes CSS.

---

## ✨ Fonctionnalités

- 🧩 Syntaxe fluide et imbriquée pour décrire des structures HTML
- 🎨 Ajout simple de classes CSS via des tableaux
- 🧱 Composants de layout : `HorizontalBar`, `VerticalBar`, `Section`, etc.
- 🖍️ Composants texte : `Titre1`, `Paragraphe`, `Span`, etc.
- 🔘 Interactions : `Bouton`, `Lien`, `Form`, `InputText`, etc.
- 🧰 Rendu final sous forme de `string`, intégrable partout

---

## 🚀 Installation

```bash
go get github.com/RafaelCoppe/Stencil-Go
```

---

## 🧑‍💻 Exemple d’utilisation

```go
package main

import (
    "fmt"
    StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
    StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
    StencilLayout "github.com/RafaelCoppe/Stencil-Go/pkg/layout"
    StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
    StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

func main() {
    html := StencilPage.Page("Exemple",
        StencilUtils.Join(
            StencilLayout.HorizontalBar([]string{
                StencilText.Paragraphe("Navigation", "text-blue-500"),
                StencilLayout.VerticalBar([]string{
                    StencilText.Titre1("Boutons", "text-xl", "font-bold"),
                    StencilLayout.HorizontalBar([]string{
                        StencilInteractions.Bouton("OK", "alert('OK')", "bg-green-500", "text-white", "px-4", "py-2"),
                        StencilInteractions.Bouton("Cancel", "alert('Cancel')", "bg-red-500", "text-white", "px-4", "py-2"),
                    }, "flex", "gap-4"),
                }),
            }, "flex", "gap-4"),
        ),
    )

    fmt.Println(html)
}
```

---

## � Architecture modulaire

Stencil-Go est organisé en packages spécialisés pour une meilleure maintenabilité :

```text
pkg/
├── form/           # Composants de formulaires
├── interactions/   # Boutons, liens, navigation
├── layout/         # Éléments de mise en page
├── media/          # Images, vidéos, audio
├── page/           # Structure de page et listes
├── text/           # Éléments textuels
├── types/          # Types et constantes
└── utils/          # Utilitaires et helpers
```

### 🔧 Importation des packages

```go
import (
    StencilForm "github.com/RafaelCoppe/Stencil-Go/pkg/form"
    StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
    StencilLayout "github.com/RafaelCoppe/Stencil-Go/pkg/layout"
    StencilMedia "github.com/RafaelCoppe/Stencil-Go/pkg/media"
    StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
    StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
    StencilTypes "github.com/RafaelCoppe/Stencil-Go/pkg/types"
    StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)
```

---

## �🔧 Fonctions disponibles

### 📄 Structure de page

| Fonction | Description |
|---------|-------------|
| `Page(title, body string, classes ...string)` | Génère une page HTML complète |
| `Div(content string, classes ...string)` | Génère un `<div>` |
| `Container(content string, classes ...string)` | Génère un conteneur centré |

### 🧱 Layout

| Fonction | Description |
|---------|-------------|
| `HorizontalBar(children []string, classes ...string)` | Flex row |
| `VerticalBar(children []string, classes ...string)` | Flex column |
| `Section(content string, classes ...string)` | Génère une section |

### ✏️ Éléments texte

| Fonction | Description |
|---------|-------------|
| `Titre1(text string, classes ...string)` | Génère un `<h1>` |
| `Titre2(text string, classes ...string)` | Génère un `<h2>` |
| `Titre3(text string, classes ...string)` | Génère un `<h3>` |
| `Titre4(text string, classes ...string)` | Génère un `<h4>` |
| `Titre5(text string, classes ...string)` | Génère un `<h5>` |
| `Titre6(text string, classes ...string)` | Génère un `<h6>` |
| `Paragraphe(text string, classes ...string)` | Génère un `<p>` |
| `Span(text string, classes ...string)` | Génère un `<span>` |
| `Code(text string, classes ...string)` | Génère un `<code>` |
| `Pre(text string, classes ...string)` | Génère un `<pre>` |
| `CodeBlock(text string, classes ...string)` | Génère un `<pre><code>` |
| `Strong(text string, classes ...string)` | Génère un `<strong>` |
| `Em(text string, classes ...string)` | Génère un `<em>` |
| `Small(text string, classes ...string)` | Génère un `<small>` |
| `Mark(text string, classes ...string)` | Génère un `<mark>` |
| `Del(text string, classes ...string)` | Génère un `<del>` |
| `Ins(text string, classes ...string)` | Génère un `<ins>` |
| `Blockquote(text string, classes ...string)` | Génère un `<blockquote>` |

### 🔘 Interactions

| Fonction | Description |
|---------|-------------|
| `Bouton(label, onClick string, classes ...string)` | Génère un bouton `<button>` |
| `BoutonSubmit(text string, classes ...string)` | Génère un bouton submit |
| `BoutonReset(text string, classes ...string)` | Génère un bouton reset |
| `BoutonDisabled(text string, classes ...string)` | Génère un bouton désactivé |
| `BoutonWithId(text, id, event string, classes ...string)` | Génère un bouton avec ID |
| `Lien(href, label string, classes ...string)` | Génère un lien `<a>` |
| `NavLink(href, text string, active bool, classes ...string)` | Lien de navigation |
| `ExternalLink(href, text string, classes ...string)` | Lien externe |
| `MailtoLink(email, text string, classes ...string)` | Lien email |
| `TelLink(phone, text string, classes ...string)` | Lien téléphone |
| `DownloadLink(href, filename, text string, classes ...string)` | Lien de téléchargement |
| `Breadcrumb(items []map[string]string, classes ...string)` | Fil d'Ariane |
| `Pagination(currentPage, totalPages int, baseURL string, classes ...string)` | Pagination |

### 📥 Formulaires

| Fonction | Description |
|---------|-------------|
| `Form(action, method string, elements []string, classes ...string)` | Génère un formulaire |
| `Input(inputType, name, placeholder string, classes ...string)` | Champ input générique |
| `InputText(name, placeholder, value string, classes ...string)` | Champ texte |
| `InputEmail(name, placeholder, value string, classes ...string)` | Champ email |
| `InputPassword(name, placeholder string, classes ...string)` | Champ mot de passe |
| `InputNumber(name, placeholder string, value, min, max int, classes ...string)` | Champ numérique |
| `InputDate(name, value string, classes ...string)` | Champ date |
| `InputFile(name, accept string, multiple bool, classes ...string)` | Champ fichier |
| `TextArea(name, placeholder, value string, rows, cols int, classes ...string)` | Zone de texte |
| `Select(name string, options map[string]string, selected string, classes ...string)` | Liste déroulante |
| `CheckBox(name, value, label string, checked bool, classes ...string)` | Case à cocher |
| `RadioButton(name, value, label string, checked bool, classes ...string)` | Bouton radio |
| `RadioGroup(name string, options []map[string]interface{}, classes ...string)` | Groupe de radio |
| `Range(name string, min, max, value int, classes ...string)` | Slider |
| `Color(name, value string, classes ...string)` | Sélecteur de couleur |
| `Hidden(name, value string)` | Champ caché |
| `Label(forInput, text string, classes ...string)` | Label |
| `FieldSet(legend, content string, classes ...string)` | Fieldset |
| `FormGroup(labelText, inputHTML string, classes ...string)` | Groupe de formulaire |

### 📋 Listes et tableaux

| Fonction | Description |
|---------|-------------|
| `Ul(items []string, classes ...string)` | Liste non ordonnée |
| `Ol(items []string, classes ...string)` | Liste ordonnée |
| `Li(content string, classes ...string)` | Élément de liste |
| `Dl(items map[string]string, classes ...string)` | Liste de définitions |
| `Table(headers []string, rows [][]string, classes ...string)` | Table complète |
| `TableSimple(rows [][]string, classes ...string)` | Table simple |
| `Tr(cells []string, classes ...string)` | Ligne de table |
| `Td(content string, classes ...string)` | Cellule de données |
| `Th(content string, classes ...string)` | Cellule d'en-tête |

### 🖼️ Médias

| Fonction | Description |
|---------|-------------|
| `Image(src, alt string, classes ...string)` | Image |
| `ImageWithTitle(src, alt, title string, classes ...string)` | Image avec titre |
| `ResponsiveImage(src, srcset, alt string, classes ...string)` | Image responsive |
| `Figure(src, alt, caption string, classes ...string)` | Figure avec légende |
| `Video(src string, controls bool, classes ...string)` | Vidéo |
| `VideoWithOptions(src string, controls, autoplay, loop, muted bool, classes ...string)` | Vidéo avec options |
| `VideoWithSources(sources map[string]string, controls bool, classes ...string)` | Vidéo multi-sources |
| `Audio(src string, controls bool, classes ...string)` | Audio |
| `Iframe(src string, width, height int, classes ...string)` | IFrame |

### 🧩 Utilitaires

| Fonction | Description |
|---------|-------------|
| `HR(classes ...string)` | Ligne horizontale |
| `Br()` | Saut de ligne |
| `Nbsp()` | Espace insécable |
| `Comment(text string)` | Commentaire HTML |
| `Join(elements ...string)` | Concatène les éléments (helper) |
| `EscapeHTML(text string)` | Échappe les caractères HTML |
| `ClassBuilder(classes map[string]bool)` | Constructeur de classes conditionnel |
| `Wrap(tag, content string, attributes map[string]string, classes ...string)` | Élément HTML générique |
| `SelfClosingTag(tag string, attributes map[string]string, classes ...string)` | Balise auto-fermante |

### 🏷️ Types

| Type | Description |
|------|-------------|
| `HTMLElement` | Structure d'élément HTML |
| `FormElement` | Structure d'élément de formulaire |
| `LinkTarget` | Cibles de liens (`_blank`, `_self`, etc.) |
| `ButtonType` | Types de boutons (button, submit, reset) |
| `InputType` | Types d'inputs (text, email, password, etc.) |
| `MediaControls` | Options de contrôles média |

---

## 🧪 Tests

Le projet inclut une suite de tests complète pour tous les packages :

```bash
# Exécuter tous les tests
go test ./...

# Tests avec coverage
go test -cover ./...

# Tests spécifiques par package
go test ./tests/form_test.go
go test ./tests/interactions_test.go
go test ./tests/layout_test.go
go test ./tests/page_test.go
go test ./tests/text_test.go
```

### 📊 Couverture de tests

Les tests couvrent :

- ✅ Tous les composants de formulaires
- ✅ Interactions (boutons, liens)
- ✅ Éléments de layout
- ✅ Structure de pages et listes
- ✅ Composants textuels
- ✅ Fonctions utilitaires

---

## 🌐 Intégration WebAssembly

Stencil-Go est conçu pour fonctionner parfaitement avec le [Framework WebAssembly Stencil](https://github.com/RafaelCoppe/Stencil-Framework) :

- **Compilation WebAssembly** : Générez des applications web performantes
- **Routage automatique** : Système de routes basé sur fichiers (comme Next.js)
- **Gestion d'état** : État réactif avec re-rendu automatique
- **Événements** : Système d'événements JavaScript intégré

### 🔗 Utilisation avec le framework

```go
// Dans une page du framework
func (p *MyPage) Render() string {
    return StencilPage.Container(
        StencilUtils.Join(
            StencilText.Titre1("Ma Page", "text-2xl", "font-bold"),
            StencilText.Paragraphe("Contenu de la page"),
            StencilInteractions.Bouton("Action", "handleAction", "btn", "btn-primary"),
        ),
        "container", "mx-auto", "p-4",
    )
}
```

---

## 🤝 Contribuer

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une *issue* ou une *pull request*. On adore les nouvelles idées ✨

### 🚀 Pour contribuer

1. Fork le projet
2. Créez une branche pour votre fonctionnalité (`git checkout -b feature/AmazingFeature`)
3. Committez vos changements (`git commit -m 'Add some AmazingFeature'`)
4. Poussez vers la branche (`git push origin feature/AmazingFeature`)
5. Ouvrez une Pull Request

### 📋 Guidelines

- Ajoutez des tests pour toute nouvelle fonctionnalité
- Respectez le style de code existant
- Documentez les nouvelles fonctions
- Vérifiez que tous les tests passent avec `go test ./...`

---

## 📄 Licence

Stencil est distribué sous licence MIT. Voir [LICENSE](./LICENSE) pour plus d'informations.
