package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error{

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s... \n", name)
	data, err := cfg.pokeapiClient.FetchPokemonData(name) 
	if err != nil {
	return err	
	}

	if rand.Float64() < 1.0 - float64(data.BaseExperience)/300.0 {
		fmt.Printf("%s was captured!\n", data.Name)
		cfg.pokedex[name] = data

	} else {
		fmt.Printf("%s escaped!\n", data.Name)
	}

	return nil
}
