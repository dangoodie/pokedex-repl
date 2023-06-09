package main

import (
	"fmt"
	"math/rand"
	"errors"
)


func commandCatch(cfg *config) error {
	pokemonPointer := cfg.value
	if pokemonPointer == nil{
		return errors.New("no value for pokemon to catch")
	}
	pokemon := *pokemonPointer
	fmt.Println("Throwing a Pokeball at ", pokemon)

	resp, err := cfg.pokeAPIClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	for _, pokedexEntry := range cfg.pokedex {
		if pokedexEntry.Name == resp.Name {
			fmt.Printf("You already have %v!\n", pokemon)
			return nil
		}
	}

	difficulty := resp.BaseExperience
	roll := rand.Intn(400)

	if roll > difficulty {
		fmt.Printf("You caught %v!\n", pokemon)
		cfg.pokedex = append(cfg.pokedex, resp)
	} else {
		fmt.Printf("Oh no! %v escaped!\n", pokemon)
	}
	return nil
}