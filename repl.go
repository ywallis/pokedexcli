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

		args := []string{}
		commandName := words[0]

		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
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
	callback    func(*Config, ...string) error
}

type Config struct {
	Previous      string
	Next          string
	pokeapiClient *pokeapi.Client
	pokedex       map[string]pokeapi.PokemonData
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch",
			description: "Try and capture a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Lists the names of all pokemons at a given location",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows details of a pokedexentry",
			callback:    commandInspect,
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
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}
