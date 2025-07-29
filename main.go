package main

import (
	"time"

	"github.com/mentalcaries/pokedexcli/internal/pokeapi"
)

func main(){
    pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
    config := &config{
        pokeApiClient: pokeClient,
    }
    startRepl(config)
}

