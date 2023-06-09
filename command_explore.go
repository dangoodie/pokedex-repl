package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	locationPointer := cfg.value
	if locationPointer == nil {
		return errors.New("no value for location to explore")
	}
	location := *locationPointer

	fmt.Println("Exploring")

	resp, err := cfg.pokeAPIClient.GetLocationPokemon(location)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon in", location, ":")
	for _, pokemonEncounter := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
