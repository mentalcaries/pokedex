package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mentalcaries/pokedexcli/internal/pokeapi"
)


type cliCommand struct {
    name string
    description string
    callback func(*config, ...string) error
}

type config struct {
    pokeApiClient pokeapi.Client
    Next *string
    Previous *string
}

func cleanInput(text string) []string {

    lowerCaseString := strings.ToLower(text)
    cleanStrings := strings.Fields(lowerCaseString)
    
    return cleanStrings
}


func startRepl(config *config){
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()

        text := cleanInput(scanner.Text())
        commandName:= text[0]
        args := []string{}

        if len(text) > 1 {
            args = text[1:]
        }

        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(config, args...)
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            fmt.Println("Unknown command")
            continue
        }

    }
}

func getCommands() map[string]cliCommand{
    
    return  map[string]cliCommand{
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
        "help": {
            name: "help",
            description: "Display a help message",
            callback: commandHelp,
        },
        "map" : {
            name: "map",
            description: "Display a list of locations",
            callback: commandMapForward,
        },
        "mapb":{
            name: "mapb",
            description: "Display previous list of locations",
            callback: commandMapBack,
        },
        "explore": {
            name: "explore",
            description: "Allows a user to see all Pokemon in an area",
            callback: commandExplore,
        },
        "catch": {
            name: "catch",
            description: "Gives user a chance to catch a pokemon",
            callback: commandCatch,
        },
    }
}