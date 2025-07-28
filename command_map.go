package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var baseUrl string = "https://pokeapi.co/api/v2/location-area/"

type locationData struct {
    Name string `json:"name"`
    URL string  `json:"url"`
}

type locationResult struct {
    Count int `json:"count"`
    Next *string  `json:"next"`
    Previous *string `json:"previous"`
    Results []locationData `json:"results"`
}

func commandMap(config *config)error {
    requestUrl := baseUrl
    if config.Next != nil {
        requestUrl = *config.Next
    }
    res, err := http.Get(requestUrl)
    if err != nil{
        fmt.Println(err)
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
    }
    
    
    if res.StatusCode > 299 {
        fmt.Printf("Response failed with status code: %d", res.StatusCode)
    }

    var locationResponse locationResult

    if err := json.Unmarshal(body, &locationResponse); err != nil {
        fmt.Printf("Error reading JSON %v", err)
    }

    
    for _, result := range locationResponse.Results {
        fmt.Println(result.Name)
    }

    config.Next= locationResponse.Next
    config.Previous = locationResponse.Previous

    return nil
}

func commandMapBack(config *config)error {

    if config.Previous == nil {
        fmt.Println("You're on the first page")
        return nil
    }
    requestUrl := *config.Previous
    res, err := http.Get(requestUrl)
    if err != nil{
        fmt.Println(err)
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
    }
    
    
    if res.StatusCode > 299 {
        fmt.Printf("Response failed with status code: %d", res.StatusCode)
    }

    var locationResponse locationResult

    if err := json.Unmarshal(body, &locationResponse); err != nil {
        fmt.Printf("Error reading JSON %v", err)
    }

    
    for _, result := range locationResponse.Results {
        fmt.Println(result.Name)
    }

    config.Next= locationResponse.Next
    config.Previous = locationResponse.Previous
    
    return nil
}