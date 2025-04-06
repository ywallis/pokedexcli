package main

import (
	"fmt"
)


func commandExplore(config *Config, args ...string) error {

	if len(args) != 1 {
		fmt.Println("missing area argument!")
		return nil
	}
	area := args[0]
	fmt.Printf("Exploring: %s\n", area)
	url := "https://pokeapi.co/api/v2/location-area/" + area
	data, err := config.pokeapiClient.FetchLocationData(url)

	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, loc := range data.Encounters {
		fmt.Println(loc.Pokemon.Name)
	}

	return nil

}
