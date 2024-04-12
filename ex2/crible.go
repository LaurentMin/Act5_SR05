package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func filtrer(i int, c chan int, resultChan chan int, wg *sync.WaitGroup) {
	
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
	wg.Done() // Marquer la fin de la goroutine
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
