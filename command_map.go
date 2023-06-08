package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeAPIClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, locationArea := range resp.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("cannot go back any further")
	}
	resp, err := cfg.pokeAPIClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, locationArea := range resp.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}
