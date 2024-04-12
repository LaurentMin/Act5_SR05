package main

import (
    "flag"
    "fmt"
)

func main() {

    // Option -n sur la ligne de commande qui attend un entier
    // (valeur par défaut : 53, message pour l'aide avec l'option -h : "nombre")
    // p_num est un pointeur sur entier.
    p_num := flag.Int("n", 53, "nombre")
    flag.Parse()

    // Print affiche selon le type fourni.
    fmt.Print("\n nombre fourni = ", *p_num, "\n")
}