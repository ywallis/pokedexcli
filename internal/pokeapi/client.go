package pokeapi

import (
	"net/http"
	"time"

	"github.com/ywallis/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, lifetime time.Duration) *Client {
	return &Client{
		cache:      *pokecache.NewCache(lifetime),
		httpClient: http.Client{Timeout: timeout},
	}
}
