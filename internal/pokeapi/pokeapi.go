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
