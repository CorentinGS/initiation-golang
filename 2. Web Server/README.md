# Web Server

## Exercices

1. Tâche : Créer un endpoint de Base

   - Description : Créez un endpoint qui répond avec un message statique (par exemple, "Bienvenue sur l'API") lorsqu'il est accédé.
   - Objectif : Familiariser les participants avec la création de points d'extrémité en Go et la gestion des requêtes HTTP.

2. Tâche : Implémenter un endpoint GET pour Récupérer des Donnée

   - Description : Créez un endpoint qui gère les requêtes GET et renvoie des données factices (par exemple, une liste d'utilisateurs ou de produits) au format JSON.
   - Objectif : Enseigner aux participants comment implémenter un endpoint en lecture seule pour récupérer des données du serveur.
   - Fonctions à utiliser: toJSON(users) -> string
   - Commande pour tester: `curl http://localhost:8080/users`

3. Tâche : Implémenter un endpoint POST pour Ajouter des Donnée

   - Description : Créez un endpoint qui gère les requêtes POST et ajoute de nouvelles données (par exemple, un nouvel utilisateur ou un nouveau produit) au serveur.
   - Objectif : Introduire aux participants la gestion des requêtes POST et l'analyse des corps de requête pour ajouter des données au serveur.
   - Commande pour tester: `curl -X POST -d "username=exampleuser" http://localhost:8080/new-users`

4. Tâche: Jouer avec les différents endpoints pour modifier les résultats et ajouter un endpoint de votre choix.