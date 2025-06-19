package pokeapi

import (
	"net/http"
	"time"

	"github.com/L-chaCon/pokedex/internal/pokecache"
)

type Client struct {
	pokeCache  pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cleanTime time.Duration) Client {
	return Client{
		pokeCache:  pokecache.NewCache(cleanTime),
		httpClient: http.Client{Timeout: timeout},
	}
}
