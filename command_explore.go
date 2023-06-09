package main 

import "fmt"

func commandExplore(cfg *config) error {
	fmt.Println("Exploring")

	location := cfg.userFields[1]
	resp, err := cfg.pokeAPIClient.GetLocationPokemon(location)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon in", location, ":")
	for _, pokemonEncounter := range resp.PokemonEncounters {
		fmt.Printf("%s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
