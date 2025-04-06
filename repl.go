package main

import (
	"bufio"
	"fmt"
	"github.com/ywallis/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

func startRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	config.Previous = ""
	config.Next = "https://pokeapi.co/api/v2/location-area/"

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		var argument string
		commandName := words[0]
		if len(words) < 2 {
			argument = ""
		} else {
			argument = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, argument)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

type Config struct {
	Previous      string
	Next          string
	pokeapiClient *pokeapi.Client
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 map locations",
			callback:    commandMapPrevious,
		},
		"explore": {
			name:        "explore",
			description: "Lists the names of all pokemons at a given location",
			callback:    commandExplore,
		},
	}
}
