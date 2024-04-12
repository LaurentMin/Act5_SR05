# SR05 Activité 5

## Exercice 1 : SOMME

### Utilisation

Compiler avec go build sum.go
Exécuter avec ./sum -n nb
avec nb un nombre entier pour faire la somme de 1 à nb (par défaut nb = 10)

### Description

Nous avons d'abord une fonction sommer() qui prend en paramètre une liste de nombres et qui envoie la somme de ces nombres dans un channel.

On construit ensuite une liste de nombres de 1 à nb (nb étant le nombre passé en paramètre du programme ou 10 par défaut).

On découpe ce tableau en nPCU sous-tableaux de taille égale avec plus ou moins 1 (nPCU étant le nombre de processeurs logiques de la machine).

On lance ensuite nPCU goroutines qui calculent la somme de chaque sous-tableau qui sont envoyer dans un channel.

On récupère les sommes partielle dans le channel et on les additionne pour obtenir la somme totale.

### Exemple

```bash
$ ./sum -n 10
La somme de 1 à 10 est 55
```

## Exercice 2 : Crible de Hoare

### Utilisation

Compiler avec go build crible.go
Exécuter avec ./crible -n nb
avec nb un nombre entier pour afficher les nombres premiers inférieurs à nb (par défaut nb = 100)