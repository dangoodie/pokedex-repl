package main

import (
	"fmt"
	"math/rand"
)


func commandCatch(cfg *config) error {
	pokemon := cfg.userFields[1]
	fmt.Println("Throwing a Pokeball at ", pokemon)

	resp, err := cfg.pokeAPIClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	for _, pokedexEntry := range cfg.pokedex {
		if pokedexEntry.Name == resp.Name {
			fmt.Printf("You already have %s!\n", pokemon)
			return nil
		}
	}

	difficulty := resp.BaseExperience
	roll := rand.Intn(400)

	if roll > difficulty {
		fmt.Printf("You caught %s!\n", pokemon)
		cfg.pokedex = append(cfg.pokedex, resp)
	} else {
		fmt.Printf("Oh no! %s escaped!\n", pokemon)
	}
	return nil
}