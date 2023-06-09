package main

import (
	"time"

	"github.com/dangoodie/pokedex-repl/internal/pokeapi"
)

type config struct {
	pokeAPIClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	userFields          []string
	pokedex             []pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeAPIClient: pokeapi.NewClient(time.Minute * 5),
	}

	startREPL(&cfg)
}
