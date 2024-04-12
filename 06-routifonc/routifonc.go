package main

import (
    "fmt"
    "time"
)

func travailler(k int) {
    fmt.Print("\n Début du long travail ", k)
    time.Sleep(1 * time.Second)
    fmt.Print("\n Fin du long travail ", k)
}

func travailler_plus(k int) {
    fmt.Print("\n   Début du long travail plus ", k)
    time.Sleep(1 * time.Second)
    fmt.Print("\n   Fin du long travail plus ", k)
}

func main () {

    for i := 0 ; i<5 ; i++ {
        go travailler(i)
        go travailler_plus(i)
    }

    fmt.Scanln() // Pour attendre la fin des goroutines...    
}