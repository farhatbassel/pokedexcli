package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInformation(pokemon string) (Pokemon, error) {
    url := baseURL + "/pokemon/" + pokemon

    if val, ok := c.cache.Get(url); ok {
        pokemonResponse := Pokemon{}
        err := json.Unmarshal(val, &pokemonResponse)
        if err != nil {
            return Pokemon{}, err
        }
        return pokemonResponse, nil
    }

    res, err := http.Get(url)

    if err != nil {
        fmt.Printf("Error occured when fetching the data %v", err)
    }

    body, err := io.ReadAll(res.Body)
    defer res.Body.Close()

    if res.StatusCode == 404 {
        return Pokemon{}, fmt.Errorf("Pokemon %v not found", pokemon)
    }
    if res.StatusCode > 299 {
        return Pokemon{}, 
        fmt.Errorf(" Response failed with status code: %d\n and\n body: %s\n", 
        res.StatusCode, body)
    }

    if err != nil {
        return Pokemon{}, fmt.Errorf("%v", err)
    }

    pokemonResp := Pokemon{}
    err = json.Unmarshal(body, &pokemonResp)

    if err != nil {
        return Pokemon{}, err
    }

    c.cache.Add(url, body)
    return pokemonResp, nil
}
