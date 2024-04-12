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
	fmt.Println("tableau =", tab)

	nbCPU := runtime.NumCPU()
	fmt.Println("nbCPU =", nbCPU)

	portion := *p_num / nbCPU
	reste := *p_num % nbCPU
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

	fmt.Println("somme =", s)
}
