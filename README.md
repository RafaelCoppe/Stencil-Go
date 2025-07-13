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
go get github.com/RafaelCoppe/stencil
```

---

## 🧑‍💻 Exemple d’utilisation

```go
package main

import (
    "fmt"
    "stencil"
)

func main() {
    html := stencil.Page("Exemple",
        []string{
            stencil.HorizontalBar([]string{
                stencil.Paragraphe("Navigation", []string{"text-blue-500"}),
                stencil.VerticalBar([]string{
                    stencil.Titre1("Boutons", []string{"text-xl", "font-bold"}),
                    stencil.HorizontalBar([]string{
                        stencil.Bouton("OK", "alert('OK')", []string{"bg-green-500", "text-white", "px-4", "py-2"}),
                        stencil.Bouton("Cancel", "alert('Cancel')", []string{"bg-red-500", "text-white", "px-4", "py-2"}),
                    }, nil),
                }, nil),
            }, []string{"flex", "gap-4"}),
        },
    )

    fmt.Println(html)
}
```

---

## 🔧 Fonctions disponibles

### 📄 Structure de page

| Fonction | Description |
|---------|-------------|
| `Page(title string, bodyElements []string)` | Génère une page HTML complète |
| `Div(children string, extraClasses ...string)` | Génère un `<div>` |
| `Container(children string, extraClasses ...string)` | Génère un conteneur centré |

### 🧱 Layout

| Fonction | Description |
|---------|-------------|
| `HorizontalBar(children []string, extraClasses ...string)` | Flex row |
| `VerticalBar(children []string, extraClasses ...string)` | Flex column |
| `Section(children []string, extraClasses ...string)` | Génère une section |
| `Article(children []string, extraClasses ...string)` | Génère un article |
| `Header(children []string, extraClasses ...string)` | Génère un header |
| `Footer(children []string, extraClasses ...string)` | Génère un footer |
| `Nav(children []string, extraClasses ...string)` | Génère une navigation |
| `Main(children []string, extraClasses ...string)` | Génère un main |

### ✏️ Éléments texte

| Fonction | Description |
|---------|-------------|
| `Titre1(text string, extraClasses ...string)` | Génère un `<h1>` |
| `Titre2(text string, extraClasses ...string)` | Génère un `<h2>` |
| `Titre3(text string, extraClasses ...string)` | Génère un `<h3>` |
| `Titre4(text string, extraClasses ...string)` | Génère un `<h4>` |
| `Titre5(text string, extraClasses ...string)` | Génère un `<h5>` |
| `Titre6(text string, extraClasses ...string)` | Génère un `<h6>` |
| `Paragraphe(text string, extraClasses ...string)` | Génère un `<p>` |
| `Span(text string, extraClasses ...string)` | Génère un `<span>` |
| `Code(text string, extraClasses ...string)` | Génère un `<code>` |
| `Pre(text string, extraClasses ...string)` | Génère un `<pre>` |
| `CodeBlock(text string, extraClasses ...string)` | Génère un `<pre><code>` |
| `Strong(text string, extraClasses ...string)` | Génère un `<strong>` |
| `Em(text string, extraClasses ...string)` | Génère un `<em>` |
| `Small(text string, extraClasses ...string)` | Génère un `<small>` |
| `Mark(text string, extraClasses ...string)` | Génère un `<mark>` |
| `Del(text string, extraClasses ...string)` | Génère un `<del>` |
| `Ins(text string, extraClasses ...string)` | Génère un `<ins>` |
| `Blockquote(text string, extraClasses ...string)` | Génère un `<blockquote>` |

### 🔘 Interactions

| Fonction | Description |
|---------|-------------|
| `Bouton(label string, onClick string, extraClasses ...string)` | Génère un bouton `<button>` |
| `BoutonSubmit(text string, extraClasses ...string)` | Génère un bouton submit |
| `BoutonReset(text string, extraClasses ...string)` | Génère un bouton reset |
| `BoutonDisabled(text string, extraClasses ...string)` | Génère un bouton désactivé |
| `BoutonWithId(text, id, event string, extraClasses ...string)` | Génère un bouton avec ID |
| `Lien(href, label string, extraClasses ...string)` | Génère un lien `<a>` |
| `NavLink(href, text string, active bool, extraClasses ...string)` | Lien de navigation |
| `ExternalLink(href, text string, extraClasses ...string)` | Lien externe |
| `MailtoLink(email, text string, extraClasses ...string)` | Lien email |
| `TelLink(phone, text string, extraClasses ...string)` | Lien téléphone |
| `DownloadLink(href, filename, text string, extraClasses ...string)` | Lien de téléchargement |
| `Breadcrumb(items []map[string]string, extraClasses ...string)` | Fil d'Ariane |
| `Pagination(currentPage, totalPages int, baseURL string, extraClasses ...string)` | Pagination |

### 📥 Formulaires

| Fonction | Description |
|---------|-------------|
| `Form(action, method string, elements []string, extraClasses ...string)` | Génère un formulaire |
| `Input(inputType, name, placeholder string, extraClasses ...string)` | Champ input générique |
| `InputText(name, placeholder, value string, extraClasses ...string)` | Champ texte |
| `InputEmail(name, placeholder, value string, extraClasses ...string)` | Champ email |
| `InputPassword(name, placeholder string, extraClasses ...string)` | Champ mot de passe |
| `InputNumber(name, placeholder string, value, min, max int, extraClasses ...string)` | Champ numérique |
| `InputDate(name, value string, extraClasses ...string)` | Champ date |
| `InputFile(name, accept string, multiple bool, extraClasses ...string)` | Champ fichier |
| `TextArea(name, placeholder, value string, rows, cols int, extraClasses ...string)` | Zone de texte |
| `Select(name string, options map[string]string, selected string, extraClasses ...string)` | Liste déroulante |
| `CheckBox(name, value, label string, checked bool, extraClasses ...string)` | Case à cocher |
| `RadioButton(name, value, label string, checked bool, extraClasses ...string)` | Bouton radio |
| `RadioGroup(name string, options []map[string]interface{}, extraClasses ...string)` | Groupe de radio |
| `Range(name string, min, max, value int, extraClasses ...string)` | Slider |
| `Color(name, value string, extraClasses ...string)` | Sélecteur de couleur |
| `Hidden(name, value string)` | Champ caché |
| `Label(forInput, text string, extraClasses ...string)` | Label |
| `FieldSet(legend, content string, extraClasses ...string)` | Fieldset |
| `FormGroup(labelText, inputHTML string, extraClasses ...string)` | Groupe de formulaire |

### 📋 Listes et tableaux

| Fonction | Description |
|---------|-------------|
| `Ul(items []string, extraClasses ...string)` | Liste non ordonnée |
| `Ol(items []string, extraClasses ...string)` | Liste ordonnée |
| `Li(content string, extraClasses ...string)` | Élément de liste |
| `Dl(items map[string]string, extraClasses ...string)` | Liste de définitions |
| `Table(headers []string, rows [][]string, extraClasses ...string)` | Table complète |
| `TableSimple(rows [][]string, extraClasses ...string)` | Table simple |
| `Tr(cells []string, extraClasses ...string)` | Ligne de table |
| `Td(content string, extraClasses ...string)` | Cellule de données |
| `Th(content string, extraClasses ...string)` | Cellule d'en-tête |

### 🖼️ Médias

| Fonction | Description |
|---------|-------------|
| `Image(src, alt string, extraClasses ...string)` | Image |
| `ImageWithTitle(src, alt, title string, extraClasses ...string)` | Image avec titre |
| `ResponsiveImage(src, srcset, alt string, extraClasses ...string)` | Image responsive |
| `Figure(src, alt, caption string, extraClasses ...string)` | Figure avec légende |
| `Video(src string, controls bool, extraClasses ...string)` | Vidéo |
| `VideoWithOptions(src string, controls, autoplay, loop, muted bool, extraClasses ...string)` | Vidéo avec options |
| `VideoWithSources(sources map[string]string, controls bool, extraClasses ...string)` | Vidéo multi-sources |
| `Audio(src string, controls bool, extraClasses ...string)` | Audio |
| `Iframe(src string, width, height int, extraClasses ...string)` | IFrame |

### 🧩 Utilitaires

| Fonction | Description |
|---------|-------------|
| `HR(extraClasses ...string)` | Ligne horizontale |
| `Br()` | Saut de ligne |
| `Nbsp()` | Espace insécable |
| `Comment(text string)` | Commentaire HTML |
| `Join(elements ...string)` | Concatène les éléments (helper) |
| `EscapeHTML(text string)` | Échappe les caractères HTML |
| `ClassBuilder(classes map[string]bool)` | Constructeur de classes conditionnel |
| `Wrap(tag, content string, attributes map[string]string, extraClasses ...string)` | Élément HTML générique |
| `SelfClosingTag(tag string, attributes map[string]string, extraClasses ...string)` | Balise auto-fermante |

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

```bash
go test ./...
```

---

## 🤝 Contribuer

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une *issue* ou une *pull request*. On adore les nouvelles idées ✨

---

## 📄 Licence

Stencil est distribué sous licence MIT. Voir [LICENSE](./LICENSE) pour plus d'informations.
