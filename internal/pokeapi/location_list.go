package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mentalcaries/pokedexcli/internal/pokecache"
)

const cacheInterval = 5 * time.Minute // 5 minute interval

var apiCache pokecache.Cache = pokecache.NewCache(cacheInterval)

func (c *Client) ListLocations(pageUrl *string) (LocationResult, error){
    
    requestUrl := baseUrl + "/location-area"
    if pageUrl!= nil {
        requestUrl = *pageUrl
    }

    locationResponse := LocationResult{}

    cachedResponse, exists := apiCache.Get(requestUrl)
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

    apiCache.Add(requestUrl, body)

    if err := json.Unmarshal(body, &locationResponse); err != nil {
        return LocationResult{}, err
    }

    return locationResponse, nil
}