package main

import "fmt"

func commandMap(config *config, area string) error {
    locationResp, err := config.pokeapiClient.ListLocations(config.nextLocationURL)

    if err != nil {
        return err
    }

    config.nextLocationURL = locationResp.Next
    config.prevLocationURL = locationResp.Previous

    for _, loc := range(locationResp.Results) {
        fmt.Println(loc.Name)
    }
    return nil
}

func commandMapBack(config *config, area string) error {
    if config.prevLocationURL == nil {
        return fmt.Errorf("You are on the first page")
    }
    locationResp, err := config.pokeapiClient.ListLocations(config.prevLocationURL)

    if err != nil {
        return err
    }

    config.nextLocationURL = locationResp.Next
    config.prevLocationURL = locationResp.Previous

    for _, loc := range(locationResp.Results) {
        fmt.Println(loc.Name)
    }
    return nil
}
