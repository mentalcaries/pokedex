package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error{

    pokemon := ""
    if len (args) < 1 {
        return fmt.Errorf("please enter a valid pokemon name")
    }

    pokemon = args[0]

    fmt.Println("Throwing a Pokeball at " + pokemon + "...")

    pokemonDetails, err := config.pokeApiClient.GetPokemonStats(pokemon)
    
    if err != nil {
        fmt.Println("No details found for this pokemon")
        return err
    }

    userExperience := rand.Intn(400)
    pokemonExperience := pokemonDetails.BaseExperience

    if userExperience < pokemonExperience {
        fmt.Println(pokemon + " got away. Try again")
        return nil
    }

    fmt.Println("You caught " + pokemon + "...")
    fmt.Println("Saving to your Pokedex...")

    Pokedex[pokemon] = pokemonDetails

    return nil
}