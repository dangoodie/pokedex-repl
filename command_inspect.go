package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config) error {
	pokemonPointer := cfg.value
	if pokemonPointer == nil {
		return errors.New("no value for pokemon to inspect")
	}
	pokemon := *pokemonPointer

	resp, ok := cfg.pokedex[pokemon]
	if !ok {
		return errors.New("you have not caught " + pokemon)
	}

	fmt.Println("Name:", resp.Name)
	fmt.Println("Height:", resp.Height)
	fmt.Println("Weight:", resp.Weight)

	fmt.Println("Stats:")
	for _, stat := range resp.Stats {
		fmt.Println("-", stat.Stat.Name, ":", stat.BaseStat )
	}

	fmt.Println("Types:")
	for _, pokemonType := range resp.Types {
		fmt.Println("-", pokemonType.Type.Name)
	}
	return nil
}
