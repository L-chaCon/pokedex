package main

import (
	"time"

	"github.com/L-chaCon/pokedex/internal/config"
	"github.com/L-chaCon/pokedex/internal/pokeapi"
	"github.com/L-chaCon/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second)
	pokeCache := pokecache.NewCache(2 * time.Minute)
	cfg := &config.Config{PokeapiClient: pokeClient, PokeCach: pokeCache}
	startRepl(cfg)
}
