package main

import (
	"farhatbassel/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second)
    config := &config{
        pokeapiClient: pokeClient,
    }
    repl(config)
}

