package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) FetchLocationAreas(url string) (Response, error) {
	body, ok := c.cache.Get(url)
	if !ok {

		resp, err := c.httpClient.Get(url)
		if err != nil {
			return Response{}, fmt.Errorf("failed to make request: %w", err)
		}
		defer resp.Body.Close()

		bodybytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return Response{}, fmt.Errorf("failed to read body: %w", err)
		}
		body = bodybytes
		c.cache.Add(url, body)
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		return Response{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}

func (c *Client) FetchLocationData(url string) (LocationArea, error) {
	body, ok := c.cache.Get(url)
	if !ok {

		resp, err := c.httpClient.Get(url)
		if err != nil {
			return LocationArea{}, fmt.Errorf("failed to make request: %w", err)
		}
		defer resp.Body.Close()

		bodybytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, fmt.Errorf("failed to read body: %w", err)
		}
		body = bodybytes
		c.cache.Add(url, body)
	}

	var data LocationArea 
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationArea{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}

func (c *Client) FetchPokemonData(name string) (PokemonData, error){
	endpoint := "https://pokeapi.co/api/v2/pokemon/"
	fullUrl := endpoint + name
	rawData, ok := c.cache.Get(fullUrl)
	if !ok {
		resp, err := c.httpClient.Get(fullUrl)
		if err != nil {
			return PokemonData{}, fmt.Errorf("failed to make request: %w", err)
		}
		defer resp.Body.Close()

		bodybytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return PokemonData{}, fmt.Errorf("failed to read body: %w", err)
		}
		rawData = bodybytes
		c.cache.Add(fullUrl, rawData)
	}
	var data PokemonData
	if err := json.Unmarshal(rawData, &data); err != nil {
		return PokemonData{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil

}
