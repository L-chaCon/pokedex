package main

import (
	"time"

	"github.com/L-chaCon/pokedex/internal/config"
	"github.com/L-chaCon/pokedex/internal/pokeapi"
	"github.com/L-chaCon/pokedex/internal/pokemons"
)

func main() {
	pokeClient := pokeapi.NewClient(10*time.Second, 5*time.Minute)
	cfg := &config.Config{
		PokeapiClient: pokeClient,
		RandomRoof:    400,
		Pokedex:       pokemons.NewPokedex(),
	}
	startRepl(cfg)
}
