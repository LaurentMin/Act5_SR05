# SR05 Activité 5

## Exercice 1 : SOMME

### Utilisation

Compiler avec `go build sum.go`\
Exécuter avec `./sum -n nb`\
avec nb un nombre entier pour faire la somme de 1 à nb (par défaut nb = 10)

### Description

Nous avons d'abord une fonction `sommer()` qui prend en paramètre une liste de nombres, un channel d'entier et un objet WaitGroup. La fonction envoie la somme de ces nombres dans un channel.

On construit ensuite une liste de nombres de 1 à `nb` (`nb` étant le nombre passé en paramètre du programme ou 10 par défaut).

On découpe ce tableau en `nPCU` sous-tableaux de taille égale avec plus ou moins 1 (`nPCU` étant le nombre de processeurs logiques de la machine).

On lance ensuite `nPCU` goroutines qui calculent la somme de chaque sous-tableau qui sont envoyer dans un channel.

On récupère les sommes partielle dans le channel et on les additionne pour obtenir la somme totale.

### Exemple

```bash
$ ./sum -n 10
La somme de 1 à 10 est 55
```

### Code

```go
package main

import (
 "flag"
 "fmt"
 "runtime"
 "sync"
)

// sommer calcule la somme des éléments d'un tableau et envoie le résultat dans un canal
func sommer(tab []int, wg *sync.WaitGroup, c chan<- int) {
 defer wg.Done() // Indiquer à WaitGroup qu'une goroutine a terminé
 //fmt.Println("tab =", tab)
 s := 0
 for _, v := range tab {
  s += v
 }
 c <- s // Envoyer le résultat dans le canal
}

func main() {
 p_num := flag.Int("n", 10, "nombre")
 flag.Parse()

 //fmt.Println("taille du tableau =", *p_num)
 tab := make([]int, *p_num)

 for i := 1; i <= *p_num; i++ {
  tab[i-1] = i
 }
 //fmt.Println("tableau =", tab)

 nbCPU := runtime.NumCPU()
 //fmt.Println("nbCPU =", nbCPU)

 portion := *p_num / nbCPU // Calculer la taille de la portion
 reste := *p_num % nbCPU // Calculer le reste
 var flag int = 0
 //fmt.Println("portion =", portion)
 //fmt.Println("reste =", reste)

 resultat := make(chan int)
 var wg sync.WaitGroup // Utiliser WaitGroup pour attendre la fin des goroutines

 for i := 0; i < nbCPU; i++ {
  debut := i * portion + flag
  fin := debut + portion
  if reste > 0 { // Ajouter un élément supplémentaire à chaque goroutine si reste > 0
   fin += 1
   reste -= 1
   flag += 1
  }

  wg.Add(1) // Indiquer à WaitGroup qu'une nouvelle goroutine démarre
  go sommer(tab[debut:fin], &wg, resultat)
 }

 go func() {
  wg.Wait() // Attendre que toutes les goroutines soient terminées
  close(resultat) // Fermer le canal une fois toutes les goroutines terminées
 }()

 s := 0
 for sum := range resultat { // Lire les résultats jusqu'à ce que le canal soit fermé
  s += sum
 }

 fmt.Println("La somme de 1 à", *p_num, "est" , s)
}

```

## Exercice 2 : Crible de Hoare

### Utilisation

Compiler avec `go build crible.go`
Exécuter avec `./crible -n nb`
avec nb un nombre entier pour afficher les nombres premiers inférieurs ou égal à `nb` (par défaut `nb` = 10)

### Description

Nous avons une fonction `filtrer()` qui prend en paramètre un entier `i`, deux channel d'entier et un objet WaitGroup. La fonction envoie l'entier dans un channel à la goroutine suivante (la crée si elle n'existe pas encore) si il n'est pas un multiple de `i`.

On initialise se processus de goroutines imbriquées avec `go filtrer(2, list_nb, resultChan, &wg)` pour filtrer les multiples de 2.\
`list_nb` est un channel qui contient les entiers à filtrer qui seront ajouter par une goroutine anonyme.\
`resultChan` est un channel qui permettra de récupérer le i de chaque goroutine.

Un 0 est ajouter à `list_nb` pour signifier la fin du processus. Si un 0 est reçu par un goroutine `filtrer()`, il envoie un 0 à la goroutine suivante, envoie son `i` dans `resultChan` et se termine.
La dernière goroutine envoie un 0 dans `resultChan` pour signifier la fin du processus.

Par construction, les nombres premiers sont reçus dans `resultChan` dans le bon ordre.

### Exemple

```bash
$ ./crible -n 10
Les nombres premiers inférieurs ou égal à 10 sont 2, 3, 5, 7
```

### Code

```go
package main

import (
 "flag"
 "fmt"
 "strconv"
 "strings"
 "sync"
)

func filtrer(i int, c chan int, resultChan chan int, wg *sync.WaitGroup) {
 defer wg.Done() // Marquer la fin de la goroutine
 var flag bool = false // Flag pour indiquer si une nouvelle goroutine a été créée
 cbis := make(chan int) // Canal pour la prochaine goroutine

 for n := range c {
  if n == 0 {
   resultChan <- i
   if !flag { 
    resultChan <- 0
    close(cbis)
   } else {
    cbis <- 0
   }
   close(c)
   break
  }

  if n%i != 0 {
   fmt.Println("goroutine ", i, " ne filtre pas ", n)
   if !flag {
    fmt.Println("goroutine ", i, " crée une nouvelle goroutine", n)
    wg.Add(1) // Ajouter une goroutine à WaitGroup
    go filtrer(n, cbis, resultChan, wg)
    flag = true
   } else {
    fmt.Println("goroutine ", i, " envoie ", n, " à la prochaine goroutine")
    cbis <- n // Envoyer le nombre à la prochaine goroutine
   }
  } else {
   fmt.Println("goroutine ", i, " filtre ", n)
  }
 }
}

func main() {
 p_num := flag.Int("n", 10, "nombre")
 flag.Parse()

 if *p_num < 2 {
  fmt.Println("Pas de nombres premiers pour n < 2")
  return
 }

 list_nb := make(chan int, *p_num)
 resultChan := make(chan int)

 var wg sync.WaitGroup // Créer un WaitGroup

 wg.Add(1) // Ajouter une goroutine à WaitGroup
 go func() {
  defer wg.Done() // Marquer la fin de la goroutine lorsque la fonction se termine
  for i := 2; i <= *p_num; i++ {
   list_nb <- i
  }
  list_nb <- 0
 }()

 wg.Add(1) // Ajouter une goroutine à WaitGroup
 go filtrer(2, list_nb, resultChan, &wg)

 var results []string

 for i := 2; i <= *p_num; i++ {
  result := <-resultChan
  if result == 0 {
   break
  }
  results = append(results, strconv.Itoa(result))
 }

 wg.Wait() // Attendre la fin de toutes les goroutines

 // Afficher les nombres premiers
 fmt.Printf("Liste des nombres premiers jusqu'à %d : %s\n", *p_num, strings.Join(results, ", "))
 close(resultChan)
}
```
