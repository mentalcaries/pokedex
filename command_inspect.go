package main

import (
	"errors"
	"fmt"
	"strconv"
)

func commandInspect(config *config, args ...string)error{

    if len(args) < 1 {
        return errors.New("please enter a valid pokemon name")
    }

    pokemon := args[0]

    pokemonDetails, exists := Pokedex[pokemon]
    if !exists {
        return errors.New("you have not caught that pokemon")
    }

    fmt.Println("Name: " + pokemonDetails.Name)
    fmt.Println("Height: " + strconv.Itoa(pokemonDetails.Height))
    fmt.Println("Weight: " + strconv.Itoa(pokemonDetails.Weight))
    
    for _, stat := range pokemonDetails.Stats {
        fmt.Println("  -" + stat.Stat.Name + ": " + strconv.Itoa(stat.BaseStat))
    }
    
    return nil
}