package main

import (
	"fmt"
)


func commandMap(config *Config, args ...string) error {

	data, err := config.pokeapiClient.FetchLocationAreas(config.Next)

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

func commandMapPrevious(config *Config, args ...string) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := config.pokeapiClient.FetchLocationAreas(config.Previous)

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
