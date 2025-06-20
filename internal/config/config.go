package config

import (
	"github.com/L-chaCon/pokedex/internal/pokeapi"
	"github.com/L-chaCon/pokedex/internal/pokemons"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	RandomRoof       int
	Pokedex          pokemons.Pokedex
}
