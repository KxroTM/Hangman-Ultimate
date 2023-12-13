# Jeu du Pendu en Golang avec Nouvelles Fonctionnalités

Le jeu du pendu est une implémentation en Golang du classique jeu où les joueurs tentent de deviner un mot en suggérant des lettres, avec un nombre limité d'erreurs autorisées. Ce jeu inclut plusieurs fonctionnalités supplémentaires pour améliorer l'expérience de jeu.

## Nouvelles Fonctionnalités

1. **--letterFile :**
   - Affiche le mot à trouver en utilisant des lettres en ASCII art pour une expérience visuelle unique.

2. **--startWith :**
   - Permet de lancer le jeu à partir d'une sauvegarde, reprenant une partie précédemment stoppée.

3. **--dictionnaireAPI :**
   - Ajoute une fonctionnalité permettant de lancer une partie avec un mot au hasard provenant d'une API externe.

4. **--spectator :**
   - Permet de rejouer une partie enregistrée en tant que spectateur, pour revivre les moments forts ou étudier les stratégies utilisées.

## Instructions

1. **Installation :**
    - Assurez-vous d'avoir Go installé sur votre machine.
    - Clonez ce dépôt avec la commande : `git clone https://ytrack.learn.ynov.com/git/ayoussef/hangman-ultimate`

2. **Exécution :**
    - Accédez au répertoire du jeu : `cd hangman-golang`
    - Compilez le jeu en utilisant la commande : `go build hangman.go`
    - Exécutez le jeu : `./hangman`

3. **Règles du jeu :**
    - Le jeu choisit un mot aléatoire parmi une liste prédéfinie, ou utilise la fonctionnalité `dictionnaireAPI` pour obtenir un mot aléatoire.
    - Le joueur propose des lettres pour deviner le mot.
    - Utilisez `letterFile` pour afficher le mot à trouver en ASCII art.
    - Utilisez `startWith` pour reprendre une partie depuis une sauvegarde.
    - Le jeu se termine lorsque le joueur devine le mot ou atteint le nombre maximum d'erreurs.

## Auteur
Ammari Youssef, Julia Gari, Elmir Elias, Ilangovane Jayanraj

---

Amusez-vous à jouer au Pendu en Golang avec ces nouvelles fonctionnalités !
