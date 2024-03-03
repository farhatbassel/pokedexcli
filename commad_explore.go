package main

import "fmt"

func commandExplore(config *config, area string) error {
    if area == "" {
        return fmt.Errorf("Please enter an area")
    }

    fmt.Printf("Exploring %v...\n", area)

    locationAreaResp, err := config.pokeapiClient.ListPokemonInLocations(area)

    if err != nil {
        return err
    }

    if len(locationAreaResp.PokemonEncounters) == 0 {
        fmt.Printf("No Pokemons were found in %v\n", area)
        return nil
    }

    fmt.Println("Found Pokemons: ")
    for _, encounters := range(locationAreaResp.PokemonEncounters) {
        fmt.Printf("- %v\n", encounters.Pokemon.Name)
    }

    return nil
}
