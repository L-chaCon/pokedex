package config

import "github.com/L-chaCon/pokedex/internal/pokeapi"

type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}
