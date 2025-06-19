package main

import (
	"time"

	"github.com/L-chaCon/pokedex/internal/config"
	"github.com/L-chaCon/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second)
	cfg := &config.Config{PokeapiClient: pokeClient}
	startRepl(cfg)
}
