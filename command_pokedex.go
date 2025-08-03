package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *config, arg ...string) error{

    if len(Pokedex) < 1 {
        return errors.New("you don't have any Pokemon. Go catch some")
    }
    
    fmt.Println("Your Pokedex:")
    for key := range Pokedex {
        fmt.Println(" - ", key)
    }

    return nil
}