package main

import (
    "fmt"
    "flag"
)

func main() {

    // Tableau de taille connue
    tabnom := [6]string{ "zéro", "un", "deux", "trois", "quatre", "cinq", }

    p_num := flag.Int("n", 5, "nombre")
    flag.Parse()

    // Tableau de taille dynamique, allant de 0 à *p_num+1
    tabnum := make([]int, *p_num + 1)


    for i:=0 ; i <= *p_num ; i++ {
        tabnum[i] = i

        if i < len(tabnom) {
            fmt.Print("\n tabnum[", i, "] = ", i, " (", tabnom[i], ")")
        }    else {
            fmt.Print("\n tabnum[", i, "] = ", i)
        }
    }

    fmt.Print("\n")
}