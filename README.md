# A la recherche de la correction perdue
Ceci est un projet d'exemple de corrigé attendu.


## Pour lancer le projet

Clonez le répertoire et allez dans le dossier courant avec votre terminal.

Pour lancer le projet si vous avec Golang sur votre ordinateur, faites simplement 
```bash
go run .
```

Si vous n'avez pas golang mais avez docker, vous pouvez lancer 
```bash
docker run -v $PWD:/app --workdir=/app --rm -it golang:1.17-alpine go run . 
```