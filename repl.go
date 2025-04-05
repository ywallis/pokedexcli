package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	commandConfig := Config{
		Previous: "",
		Next: "https://pokeapi.co/api/v2/location-area/",
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands(&commandConfig)[commandName]
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

func getCommands(config *Config) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(config) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func() error { return commandExit(config) },
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 map locations",
			callback:    func() error { return commandMap(config) },
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 map locations",
			callback:    func() error { return commandMapPrevious(config) },
		},
	}
}
