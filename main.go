package main

import (
	"farhatbassel/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
    config := &config{
        pokeapiClient: pokeClient,
        caughtPokemons: make(map[string]pokeapi.Pokemon),
    }
    repl(config)
}

