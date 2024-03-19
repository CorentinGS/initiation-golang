# Découverte de Golang

## Afficher dans la console

Pour afficher dans la console, il suffit d'importer le package `fmt` et d'utiliser les fonctions `Println` ou `Printf`

```go
package main

import "fmt"

fmt.Println("Hello World")
```

## Completer les fonctions

L'objectif de l'exercice est de completer les fonctions (`add`, `sub`, `mul`, etc) afin de se familiariser avec le langage.

### Lancer le code

Pour lancer le code, il suffit d'utiliser la commande : `go run .`

La fonction `main` sera alors exécuté.

## Tester les fonctions

Dans le fichier `main_test.go`, des tests unitaires ont été écrit afin de vérifier le bon fonctionnement des fonctions.
Pour vérifier les fonctions il suffit de lancer la commande : `go test .`

S'il n'y a pas d'erreurs, vous pouvez passer au deuxième exercice.
