package pokeapi 

type Location struct{
	Name string `json:"name"`
	Url string `json:"url"`
}

type Response struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Location `json:"results"`
}

type LocationArea struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Encounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct{
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct{
	Name string `json:"name"`
	Url string `json:"url"`
}
