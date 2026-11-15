# Portfolio Théo Sugier

Mon portfolio pour le projet Blog portfolio d'Ynov Aix, pour présenter mon parcours et mes projets.

Il y a deux versions dans le dépôt :
- la version statique en HTML/CSS (partie 1), en ligne sur GitHub Pages
- la version avec un serveur Go (partie 2), où les projets sont chargés depuis le fichier data/projects.json

Version en ligne : https://theosugier.github.io/portfolio/

## Lancer la version statique

Cloner le dépôt, puis ouvrir index.html dans un navigateur.

git clone https://github.com/theosugier/portfolio.git

## Lancer le serveur Go

Il faut Go 1.22 ou plus. Dans le dossier du projet :

go run .

Puis ouvrir http://localhost:8080 dans le navigateur. Ctrl + C pour arrêter le serveur.

## Utiliser

Le menu en haut de chaque page mène à l'accueil et aux projets. Sur la page Projets, le bouton "Consulter le projet" ouvre la page détaillée de chaque projet. Pour ajouter un projet, il suffit de l'ajouter dans data/projects.json puis de relancer le serveur.
