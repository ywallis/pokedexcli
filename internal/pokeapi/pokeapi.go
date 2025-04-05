package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/ywallis/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

func FetchLocationAreas(url string, cache *pokecache.Cache) (Response, error) {
	body, ok := cache.Get(url)
	if !ok {

		resp, err := http.Get(url)
		if err != nil {
			return Response{}, fmt.Errorf("failed to make request: %w", err)
		}
		defer resp.Body.Close()

		bodybytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return Response{}, fmt.Errorf("failed to read body: %w", err)
		}
		body = bodybytes
		cache.Add(url, body)
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		return Response{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}
