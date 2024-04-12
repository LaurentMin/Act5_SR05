package main

import (
    "fmt"
    "time"
)

func travailler(k int, c chan<- int) {
    fmt.Print("\n Début du long travail ", k)
    time.Sleep(1 * time.Second)
    fmt.Print("\n Fin du long travail ", k)
    c <- k
}

func travailler_plus(k int, c <-chan int) {
    fmt.Print("\n   Attente fin du long travail ", k)
    <- c
    fmt.Print("\n   Début du long travail bis ", k)
    time.Sleep(1 * time.Second)
    fmt.Print("\n   Fin du long travail bis ", k)
}

func main () {

    var tabchan[5] chan int;

    for i := 0 ; i<5 ; i++ {
        tabchan[i] = make(chan int)
        go travailler(i, tabchan[i])
        go travailler_plus(i, tabchan[i])
    }

    fmt.Scanln() // Pour attendre la fin des goroutines...    
}