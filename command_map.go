package main

import (
	"github.com/ywallis/pokedexcli/internal/pokeapi"
	"fmt"
	"github.com/ywallis/pokedexcli/internal/pokecache"
)


func commandMap(config *Config, cache *pokecache.Cache) error {

	data, err := pokeapi.FetchLocationAreas(config.Next, cache)

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

func commandMapPrevious(config *Config, cache *pokecache.Cache) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := pokeapi.FetchLocationAreas(config.Previous, cache)

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
