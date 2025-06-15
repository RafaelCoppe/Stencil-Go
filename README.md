# 🧱 Stencil - Framework WebAssembly déclaratif en Go

Stencil est un framework open-source en Go qui compile en WebAssembly pour générer dynamiquement du HTML dans le navigateur. Il permet de construire des interfaces utilisateur déclaratives, inspirées de SwiftUI, avec une syntaxe fluide et structurée. Plus besoin de JavaScript ni de frameworks front-end lourds.
✨ Fonctionnalités principales

- ⚙️ Compilation Go → WebAssembly (via TinyGo ou Go officiel)

- 🧩 Syntaxe déclarative et imbriquée pour le HTML

- 🎨 Gestion simplifiée des classes CSS avec des slices

- 🔄 Mise à jour dynamique du DOM (expérimental)

- 🖍️ Composants préfabriqués : Bouton, Form, Titre1, etc.

- 🚀 Autonome, sans dépendance front-end externe

- 🌐 Intégrable dans tout projet web via une simple balise "\<script>"

## 🚀 Installation & Compilation WebAssembly

1. Installer TinyGo (recommandé pour WASM)

```bash
brew install tinygo 
```

ou suivre les instructions : <https://tinygo.org/getting-started/>

2. Compiler en WebAssembly

```bash
tinygo build -o main.wasm -target wasm ./main.go
```

3. Ajouter à ton HTML

<script src="wasm_exec.js"></script>
<script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });
</script>

## 🧑‍💻 Exemple d’utilisation (Go → WASM)

```go
    package main

    import (
        "syscall/js"
        "stencil"
    )

    func main() {
        dom := stencil.Page("Exemple",
            []string{
                stencil.HorizontalBar([]string{
                    stencil.Titre1("Bienvenue sur Stencil", []string{"text-3xl", "font-bold"}),
                    stencil.Bouton("Clique-moi", "alert('Hello from WASM!')", []string{"bg-blue-500", "text-white", "px-4", "py-2"}),
                }, []string{"gap-4"}),
            },
        )

        js.Global().Get("document").Call("getElementById", "app").Set("innerHTML", dom)
        select {} // empêche la sortie du programme
    }
```

```html
<!-- index.html -->
<body>
  <div id="app"></div>
  <!-- script WebAssembly ici -->
</body>
```

## 🧰 API des composants

(identique à la version bibliothèque, mais désormais rendue côté navigateur en WASM)

Consulte les fonctions disponibles dans la documentation complète (à créer).

## 🔬 Tests

```bash
go test ./...
```

## 🤝 Contribuer

Contributions, idées, tests bienvenus ! Ouvre une issue ou une pull request pour participer 🛠️

## 📄 Licence

Stencil est distribué sous licence MIT. Voir LICENSE.
