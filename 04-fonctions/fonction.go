package main

import (
    "fmt"
    "flag"
)

func afficher(tabint []int) {
    fmt.Print("Dans afficher(), tabint =", tabint)
    fmt.Print("\n")
    for index, valeur := range (tabint) {
        fmt.Print(" tabint[", index, "] = ", valeur, "\n")
    }
}

func main() {

    tabnom := [6]string{ "zéro", "un", "deux", "trois", "quatre", "cinq", }

    p_num := flag.Int("n", 10, "nombre")
    flag.Parse()

    tabnum := make([]int, *p_num + 1)
    for i:=0 ; i <= *p_num ; i++ {
        tabnum[i] = i
        if i < 6 {
            fmt.Print("\n tabnum[", i, "] = ", i, " (", tabnom[i], ")")
        }    else {
            fmt.Print("\n tabnum[", i, "] = ", i)
        }
    }

    fmt.Print("\n")
    fmt.Print("Dans main(), tabnum =", tabnum, "\n")

    afficher(tabnum[4:8])
}