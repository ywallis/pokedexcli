package main

import (
	"time"

	"github.com/ywallis/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
