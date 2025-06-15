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
| `Div(children []string, extraClasses []string)` | Génère un `<div>` |
| `HorizontalBar(children []string, extraClasses []string)` | Flex row |
| `VerticalBar(children []string, extraClasses []string)` | Flex column |
| `Section(children []string, extraClasses []string)` | Génère une section |
| `Container(children []string, extraClasses []string)` | Génère un conteneur centré |

### ✏️ Éléments texte

| Fonction | Description |
|---------|-------------|
| `Titre1(text string, extraClasses []string)` | Génère un `<h1>` |
| `Titre2(text string, extraClasses []string)` | Génère un `<h2>` |
| `Paragraphe(text string, extraClasses []string)` | Génère un `<p>` |
| `Span(text string, extraClasses []string)` | Génère un `<span>` |

### 🔘 Interactions

| Fonction | Description |
|---------|-------------|
| `Bouton(label string, onClick string, extraClasses []string)` | Génère un bouton `<button>` |
| `Lien(label string, href string, extraClasses []string)` | Génère un lien `<a>` |

### 📥 Formulaires

| Fonction | Description |
|---------|-------------|
| `Form(children []string, action string, method string, extraClasses []string)` | Génère un formulaire |
| `InputText(name string, placeholder string, value string, extraClasses []string)` | Champ texte |
| `TextArea(name string, placeholder string, value string, extraClasses []string)` | Zone de texte |
| `CheckBox(name string, checked bool, label string, extraClasses []string)` | Case à cocher |

### 🖼️ Médias

| Fonction | Description |
|---------|-------------|
| `Image(src string, alt string, extraClasses []string)` | Image |
| `Video(src string, controls bool, extraClasses []string)` | Vidéo |

### 🧩 Divers

| Fonction | Description |
|---------|-------------|
| `HR(extraClasses []string)` | Ligne horizontale |
| `Br()` | Saut de ligne |
| `Join(elements []string)` | Concatène les éléments (helper) |

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
