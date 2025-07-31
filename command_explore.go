package main

import "fmt"

func commandExplore(config *config, args ...string) error {

    cityArea := ""
    if len(args) < 1 {
        return fmt.Errorf("invalid area")
    }
    cityArea = args[0]
    fmt.Println("Exploring " + cityArea +"...")

    cityAreaResponse, err := config.pokeApiClient.ListPokemon(cityArea)
    if err != nil {
        fmt.Println("No Pokemon found in this area")
        return err
    }

    if len (cityAreaResponse.PokemonEncounters) > 0 {

        fmt.Println("Found Pokemon:")
    }
    for _, pokemon := range cityAreaResponse.PokemonEncounters {
        fmt.Println(" - " + pokemon.Pokemon.Name)
    }

    return  nil
}