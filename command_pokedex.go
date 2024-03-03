package main

import (
	"fmt"
	"github.com/fatih/color"
)

func commandPokedex(config *config, arg string) error {
    if len(config.caughtPokemons) == 0 {
        fmt.Println("No pokemons caught yet, use " + color.CyanString("catch"))
        return nil
    }

    fmt.Printf(`
Your Pokedex:`)
    for pokemon := range(config.caughtPokemons) {
        fmt.Printf(`
    -%v`,pokemon)
    }
    fmt.Println("")
    return nil
}
