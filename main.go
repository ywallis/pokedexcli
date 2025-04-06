package main

import (
	"time"

	"github.com/ywallis/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokedex := map[string]pokeapi.PokemonData{}
	cfg := &Config{
		pokeapiClient: pokeClient,
		pokedex: pokedex,
	}
	startRepl(cfg)
}
