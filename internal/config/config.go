package config

import (
	"github.com/L-chaCon/pokedex/internal/pokeapi"
	"github.com/L-chaCon/pokedex/internal/pokecache"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	PokeCach         *pokecache.Cache
}
