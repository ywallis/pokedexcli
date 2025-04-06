package main

import (
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {

	if len(args) != 1 {
		fmt.Println("missing pokemon argument")
		return nil
	}
	name := args[0]

	data, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %d\n", data.Height)
	fmt.Printf("Weight: %d\n", data.Weight)
	fmt.Println("Stats:")
	for _, stat := range data.Stats{
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range data.Types{

		fmt.Printf("  - %s", val.Type.Name)
	}
	return nil
}
