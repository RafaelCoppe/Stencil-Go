# 🧑‍💻 Commandes principales de Stencil-Go

Ce document présente les commandes principales à utiliser dans le projet **Stencil-Go**.

---

## 🚀 Lancer les tests

### Lancer tous les tests

Pour exécuter tous les tests du projet, utilise la commande suivante :

```bash
go test ./...
```

Cette commande parcourt tous les dossiers et sous-dossiers du projet et exécute tous les tests disponibles. C'est la méthode la plus simple pour tester l'intégralité du projet.

### Lancer des tests spécifiques

Si tu veux exécuter des tests pour un fichier spécifique, comme `page_test.go`, tu peux utiliser la commande suivante :

```bash
go test ./tests/page_test.go
```

Cela exécutera uniquement les tests du fichier `page_test.go`.

### Afficher les détails des tests

Si tu veux voir les résultats détaillés des tests, utilise l'option `-v` :

```bash
go test -v ./...
```

Cela affichera chaque test avec son résultat (succès ou échec) et, en cas d'échec, l'erreur associée.

---

## 🔧 Autres commandes utiles

### Vérifier les dépendances

Pour vérifier et mettre à jour les dépendances dans le projet, utilise la commande suivante :

```bash
go mod tidy
```

Cela ajoutera les modules manquants et supprimera ceux qui ne sont plus nécessaires.

### Mettre à jour une dépendance spécifique

Pour mettre à jour une dépendance à une version spécifique, utilise :

```bash
go get <module>@<version>
```

Par exemple, pour mettre à jour un module à la version `v1.2.3` :

```bash
go get github.com/some/module@v1.2.3
```

### Exécuter le projet

Pour exécuter ton application Go, utilise la commande :

```bash
go run .
```

Cela exécutera le fichier Go principal (habituellement `main.go`) dans le répertoire actuel.

### Vérifier la version de Go

Pour vérifier la version de Go utilisée dans le projet, exécute la commande suivante :

```bash
go version
```

Cela te donnera la version de Go que tu utilises, utile pour vérifier la compatibilité.

---

## 🧪 Lancer les tests avec une couverture

Si tu veux obtenir un rapport de couverture de code lors de l'exécution des tests, utilise cette commande :

```bash
go test -cover ./...
```

Cela ajoutera des informations sur la couverture de test pour chaque fichier testé.

---

## 📦 Créer un fichier d'exécution (build)

Pour compiler le projet et créer un fichier exécutable, utilise :

```bash
go build -o stencil
```

Cela générera un fichier exécutable nommé `stencil` dans le répertoire courant.

---

## 📚 Obtenir de la documentation sur le code

Si tu souhaites obtenir de la documentation sur le package ou une fonction spécifique, tu peux utiliser la commande suivante :

```bash
go doc <package>.<function>
```

Par exemple, pour obtenir la documentation sur la fonction `Container` dans le package `page`, utilise :

```bash
go doc page.Container
```

---

## 💡 Aide Go

Si tu as besoin d'aide ou de plus d'informations sur une commande Go spécifique, utilise la commande `go help` :

```bash
go help
```

Cela affichera un résumé de toutes les commandes disponibles et de leur utilisation.

---

### Conclusion

Ce fichier contient les commandes principales à utiliser dans le projet **Stencil-Go** pour effectuer des tests, gérer les dépendances et compiler l'application. N'hésite pas à consulter la documentation Go officielle pour plus de détails.