package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mentalcaries/pokedexcli/internal/pokecache"
)


func (c *Client) GetPokemonStats(name string) (Pokemon, error) {

    const baseUrl string = "https://pokeapi.co/api/v2/pokemon/"
    requestUrl := baseUrl + name

    pokemonDetailResponse := Pokemon{}

    cachedResponse, exists := pokecache.ApiCache.Get(name)
    if exists {
        fmt.Println("Using cached data: ")
        err := json.Unmarshal(cachedResponse, &pokemonDetailResponse)
        if err != nil {
            return Pokemon{}, err
        }
        return pokemonDetailResponse, nil
    }

    req, err := http.NewRequest("GET", requestUrl, nil)
    if err != nil {
        return Pokemon{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return Pokemon{}, err
    }

    defer res.Body.Close()

    if res.StatusCode > 299 {
        return Pokemon{}, errors.New("invalid Pokemon")
    }

    body, err := io.ReadAll(res.Body)

    if err != nil {
        return Pokemon{}, err
    }


    pokecache.ApiCache.Add(name, body)

    err = json.Unmarshal(body, &pokemonDetailResponse)
    if err != nil {
        return Pokemon{}, err
    }

    return pokemonDetailResponse, nil
}