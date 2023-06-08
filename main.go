package main

import (
	"github.com/dangoodie/pokedex-repl/internal/pokeapi"
)

func main() {
	pokeAPIClient := pokeapi.NewClient()
	resp, err := pokeAPIClient.ListLocationAreas()
	if err != nil {
		panic(err)
	}
	for _, locationArea := range resp.Results {
		println(locationArea.Name)
	}
	
	// startREPL()
}
