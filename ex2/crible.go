package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func filtrer(i int, c chan int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Marquer la fin de la goroutine lorsque la fonction se termine

	var flag bool = false // Flag pour indiquer si une nouvelle goroutine a été créée
	cbis := make(chan int) // Canal pour la prochaine goroutine
	defer close(c)
	defer func() { resultChan <- i }() // Envoyer le numéro d'identification de la goroutine à la fin de la fonction

	for n := range c {
		if n == 0 {
			if !flag {
				resultChan <- 0
			} else {
				cbis <- 0
			}
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

	list_nb := make(chan int)
	resultChan := make(chan int)

	var wg sync.WaitGroup // Créer un WaitGroup

	wg.Add(1) // Ajouter une goroutine à WaitGroup
	go func() {
		defer close(list_nb)
		for i := 2; i <= *p_num; i++ {
			list_nb <- i
		}
		list_nb <- 0
	}()

	wg.Add(1) // Ajouter une goroutine à WaitGroup
	go filtrer(2, list_nb, resultChan, &wg)

	wg.Wait() // Attendre la fin de toutes les goroutines

	// Afficher les nombres premiers
	var results []string

	for i := 2; i <= *p_num; i++ {
		result := <-resultChan
		if result == 0 {
			break
		}
		results = append(results, strconv.Itoa(result))
	}

	fmt.Printf("Liste des nombres premiers jusqu'à %d : %s\n", *p_num, strings.Join(results, ", "))
	close(resultChan)
}
