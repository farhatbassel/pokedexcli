package main

import "fmt"

func commandInspect(config *config, pokemon string) error {
    pokemonInfo, ok := config.caughtPokemons[pokemon]
    if !ok {
        fmt.Printf("You have yet to catch %v\n", pokemon)
        return nil
    }
    
    fmt.Printf(`
Height: %d
Weight: %d`, 
    pokemonInfo.Height, 
    pokemonInfo.Weight)

    if len(pokemonInfo.Stats) > 0 {
        fmt.Printf(`
Stats:`)
    }
    for _, stat := range pokemonInfo.Stats {
        fmt.Printf(`
    -%v: %d`, stat.Stat.Name, stat.BaseStat)
    }

    if len(pokemonInfo.Types) > 0 {
        fmt.Printf(`
Types:`)
    }
    for _, pokemonTypes := range pokemonInfo.Types {
        fmt.Printf(`
    -%v`, pokemonTypes.Type.Name)
    }
    
    fmt.Println("")
    return nil
}
