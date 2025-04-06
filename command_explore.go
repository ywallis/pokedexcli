package main

import (
	"fmt"
)


func commandExplore(config *Config, area string) error {

	if area == "" {
		fmt.Println("missing area argument!")
		return nil
	}
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
