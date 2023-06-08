package main

import (
	"time"

	"github.com/dangoodie/pokedex-repl/internal/pokeapi"
)

type config struct {
	pokeAPIClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeAPIClient: pokeapi.NewClient(time.Hour),
	}

	startREPL(&cfg)
}
