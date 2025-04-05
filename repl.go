package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/ywallis/pokedexcli/internal/pokecache"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	commandConfig := Config{
		Previous: "",
		Next: "https://pokeapi.co/api/v2/location-area/",
	}
	cache := pokecache.NewCache(time.Duration(15) * time.Second)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands(&commandConfig, cache)[commandName]
		if exists {
			err := command.callback()
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
	callback    func() error
}

type Config struct {
	Previous string
	Next     string
}

func getCommands(config *Config, cache *pokecache.Cache) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(config, cache) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func() error { return commandExit(config) },
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 map locations",
			callback:    func() error { return commandMap(config, cache) },
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 map locations",
			callback:    func() error { return commandMapPrevious(config, cache) },
		},
	}
}
