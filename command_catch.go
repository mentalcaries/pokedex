package main

import "fmt"

func commandCatch(config *config, args ...string) error{

    pokemon := ""
    if len (args) < 1 {
        return fmt.Errorf("please enter a valid pokemon name")
    }

    pokemon = args[0]

    fmt.Println("Throwing a Pokeball at " + pokemon + "...")

    return nil
}