package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jh318/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	nextURL     *string
	previousURL *string
	pokeClient  pokeapi.Client
}

func startRepl() {
	cfg := &config{
		pokeClient: pokeapi.NewClient(5*time.Second, 5*time.Second),
	}
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputFormatted := cleanInput(input)
		if len(inputFormatted) == 0 {
			continue
		}
		cmdName := inputFormatted[0]
		args := inputFormatted[1:]
		cmd, ok := commands[cmdName]
		if ok {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List the pokemon in a location area",
			callback:    commandExplore,
		},
	}
}
