package main

import "fmt"

func commandHelp(config *config)error {
    fmt.Printf(`Welcome to the Pokedex!
Usage:

`)
    for _, commandData := range getCommands() {
        fmt.Println(commandData.name + ": " + commandData.description)
    }
    return nil
}