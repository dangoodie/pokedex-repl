package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex REPL > ")
		reader.Scan()

		userInputRaw := reader.Text()
		userInputCleaned := strings.TrimSpace(strings.ToLower(userInputRaw))
		userFields := strings.Fields(userInputCleaned)
		if len(userFields) == 0 {
			continue
		}

		cfg.userFields = userFields
		userCommand := userFields[0]

		command, ok := getCommands()[userCommand]
		if !ok {
			fmt.Printf("Unknown command: %s\n", userInputRaw)
			continue
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays 20 locations. Subsequent calls will display the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemon that can be found in a given area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon.",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows the pokemon that you have caught.",
			callback:    commandPokedex,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the application.",
			callback:    commandExit,
		},
	}
}
