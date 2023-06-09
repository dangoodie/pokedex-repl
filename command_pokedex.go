package main

import "fmt"

func commandPokedex(cfg *config) error {
	fmt.Println("Pokedex entries:")
	for _, pokedexEntry := range cfg.pokedex {
		fmt.Println("- ", pokedexEntry.Name)
	}

	return nil
}