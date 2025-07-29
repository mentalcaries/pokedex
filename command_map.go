package main

import (
	"errors"
	"fmt"
)


func commandMapForward(config *config) error {

    locationResponse, err := config.pokeApiClient.ListLocations(config.Next)
    if err != nil {
        return err
    }

    config.Next= locationResponse.Next
    config.Previous = locationResponse.Previous

    for _, result := range locationResponse.Results {
        fmt.Println(result.Name)
    }

    return nil
}

func commandMapBack(config *config)error {

    if config.Previous == nil {
        return errors.New("you're on the first page")
    }

    locationResponse, err := config.pokeApiClient.ListLocations(config.Previous)
    if err != nil {
        return err
    }

    config.Next= locationResponse.Next
    config.Previous = locationResponse.Previous
    
    for _, result := range locationResponse.Results {
        fmt.Println(result.Name)
    }

    return nil
}