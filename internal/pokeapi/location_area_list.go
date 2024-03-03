package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonInLocations(area string) (PokeApiLocationAreaResponse, error) {
    url := baseURL + "/location-area/"
    url += area

    if val, ok := c.cache.Get(url); ok {
        locationAreaResponse := PokeApiLocationAreaResponse{}
        err := json.Unmarshal(val, &locationAreaResponse)
        if err != nil {
            return PokeApiLocationAreaResponse{}, err
        }
        return locationAreaResponse, nil
    }

    res, err := http.Get(url)

    if err != nil {
        fmt.Printf("Error occured when fetching the data %v", err)
    }

    body, err := io.ReadAll(res.Body)
    defer res.Body.Close()

    if res.StatusCode == 404 {
        return PokeApiLocationAreaResponse{},
        fmt.Errorf("No pokemon found in the area %v", area)
    }
    if res.StatusCode > 299 {
        return PokeApiLocationAreaResponse{},
        fmt.Errorf("Response failed with status code: %d \n and\n body: %s\n", 
        res.StatusCode, body)
    }

    if err != nil {
        return PokeApiLocationAreaResponse{}, fmt.Errorf("%v", err)
    }

    locationAreaResp := PokeApiLocationAreaResponse{}
    err = json.Unmarshal(body, &locationAreaResp)

    if err != nil {
        return PokeApiLocationAreaResponse{}, err
    }

    c.cache.Add(url, body)
    return locationAreaResp, nil
}
