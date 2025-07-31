package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mentalcaries/pokedexcli/internal/pokecache"
)



func (c *Client) ListLocations(pageUrl *string) (LocationResult, error){
    
    requestUrl := baseUrl + "/location-area"
    if pageUrl!= nil {
        requestUrl = *pageUrl
    }

    locationResponse := LocationResult{}

    cachedResponse, exists := pokecache.ApiCache.Get(requestUrl)
    if exists {
        fmt.Println("Using Cached Data: ")
        if err := json.Unmarshal(cachedResponse, &locationResponse); err != nil {
                return LocationResult{}, err
        }

    return locationResponse, nil
    }

    req, err := http.NewRequest("GET", requestUrl, nil)
    if err != nil{
        return LocationResult{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return LocationResult{}, err
    }

    defer res.Body.Close()


    body, err := io.ReadAll(res.Body)
    if err != nil {
        return LocationResult{}, err
    }
    
    if res.StatusCode > 299 {
        fmt.Printf("Response failed with status code: %d", res.StatusCode)
    }

    pokecache.ApiCache.Add(requestUrl, body)

    if err := json.Unmarshal(body, &locationResponse); err != nil {
        return LocationResult{}, err
    }

    return locationResponse, nil
}