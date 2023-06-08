package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex REPL > ")
		reader.Scan()

		userInput := reader.Text()
		userInput = strings.TrimSpace(strings.ToLower(userInput))
		
		command, ok := getCommands()[userInput]
		if !ok {
			fmt.Printf("Unknown command: %s\n", userInput)
			continue
		} else {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the application",
			callback: commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations. Subsequent calls will display the next 20 locations.",
			callback: commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations.",
			callback: commandMapb,
		},
	}
}
