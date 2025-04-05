package main

import (
	"github.com/ywallis/pokedexcli/internal/pokeapi"
	"fmt"
)


func commandMap(config *Config) error {

	data, err := pokeapi.FetchLocationAreas(config.Next)

	config.Next = data.Next
	config.Previous = data.Previous
	if err != nil {
		return err
	}
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}

	return nil

}

func commandMapPrevious(config *Config) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := pokeapi.FetchLocationAreas(config.Previous)

	config.Next = data.Next
	config.Previous = data.Previous
	if err != nil {
		return err
	}
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}

	return nil

}
