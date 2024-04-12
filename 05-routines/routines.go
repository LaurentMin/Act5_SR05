package main

import (
    "fmt"
    "time"
    "strconv"
)

func main () {

    for i := 0 ; i<5 ; i++ {
        go func(k int) {
            fmt.Print("\n Début du long travail ", strconv.Itoa(k))
            time.Sleep(1 * time.Second)
            fmt.Print("\n Fin du long travail ", strconv.Itoa(k))
        }(i)
    }

    // Boucle infinie peu gourmande pour attendre la fin des go-routines
    for {
        time.Sleep(time.Duration(60) * time.Second)
    }
}