package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mentalcaries/pokedexcli/internal/pokecache"
)
	


func (c *Client) ListPokemon(cityArea string) (ExploreResult, error) {
    baseUrl := "https://pokeapi.co/api/v2/location-area/"

    requestUrl := baseUrl + cityArea

    exploreResponse := ExploreResult{}

    cachedResponse, exists := pokecache.ApiCache.Get(cityArea)
    if exists {
        fmt.Println("Using cached Data: ")
        err := json.Unmarshal(cachedResponse, &exploreResponse)
        if err != nil {
            return ExploreResult{}, err
        }
        return exploreResponse, nil
    }

    req, err := http.NewRequest("GET", requestUrl, nil)
    if err != nil {
        return ExploreResult{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return ExploreResult{}, err
    }

    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)

    if err != nil {
        return ExploreResult{}, nil
    }

    if res.StatusCode > 299 {
        fmt.Println("Invalid city")
        return ExploreResult{}, nil
    }

    pokecache.ApiCache.Add(cityArea, body)

    if err := json.Unmarshal(body, &exploreResponse); err != nil {
        return ExploreResult{}, nil
    }

    
    return exploreResponse, nil
}