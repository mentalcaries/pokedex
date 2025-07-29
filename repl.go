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
    callback func(*config) error
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

        
        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(config)
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
    }
}