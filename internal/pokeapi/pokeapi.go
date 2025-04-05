package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

func FetchLocationAreas(url string) (Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Response{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, fmt.Errorf("failed to read body: %w", err)
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		return Response{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}

