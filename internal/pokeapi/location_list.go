package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokeApiLocationResponse, error) {
    url := baseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }

    if val, ok := c.cache.Get(url); ok {
        locationResponse := PokeApiLocationResponse{}
        err := json.Unmarshal(val, &locationResponse)
        if err != nil {
            return PokeApiLocationResponse{}, err
        }
        return locationResponse, nil
    }

    res, err := http.Get(url)

    if err != nil {
        fmt.Printf("Error occured when fetching the data %v", err)
    }

    body, err := io.ReadAll(res.Body)
    defer res.Body.Close()

    if res.StatusCode > 299 {
        return PokeApiLocationResponse{}, fmt.Errorf(`
        Response failed with status code: %d and\nbody: %s\n`, 
        res.StatusCode, body)
    }

    if err != nil {
        return PokeApiLocationResponse{}, fmt.Errorf("%v", err)
    }

    locationResp := PokeApiLocationResponse{}
    err = json.Unmarshal(body, &locationResp)

    if err != nil {
        return PokeApiLocationResponse{}, err
    }

    c.cache.Add(url, body)
    return locationResp, nil
}
