package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

func capture(baseExperience int) bool {
    random := rand.Int() * baseExperience
    return random > 10
}

func commandCatch(config *config, pokemon string) error {
    pokemonInfo, err := config.pokeapiClient.GetPokemonInformation(pokemon)
    if err != nil {
        return err
    }
    _, ok := config.caughtPokemons[pokemon]

    if ok {
        fmt.Printf("%v already in Pokedex\n", pokemon)
        return nil
    }

    fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

    isCaught := capture(pokemonInfo.BaseExperience)

    if isCaught {
        config.caughtPokemons[pokemon] = pokemonInfo
        fmt.Printf("%v " + color.GreenString("caught") + "!\n", pokemon)
        fmt.Printf("You may now inspect it with the " + color.HiRedString("inspect") + " command\n")
        return nil
    }

    fmt.Printf("%v " + color.RedString("escaped") + "!\n", pokemon)

    return nil
}
