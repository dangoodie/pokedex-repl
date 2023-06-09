package main

import (
	"time"

	"github.com/dangoodie/pokedex-repl/internal/pokeapi"
)

type config struct {
	pokeAPIClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	command             *string
	value               *string
	pokedex             map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeAPIClient: pokeapi.NewClient(time.Minute * 5),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startREPL(&cfg)
}
