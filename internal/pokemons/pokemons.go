package pokemons

import (
	"fmt"

	"github.com/L-chaCon/pokedex/internal/pokeapi"
)

type Pokedex struct {
	pokemons map[string]pokeapi.Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{make(map[string]pokeapi.Pokemon)}
}

func (p *Pokedex) Add(name string, pokemon pokeapi.Pokemon) {
	p.pokemons[name] = pokemon
}

func (p *Pokedex) PrintPokemons() {
	fmt.Println("This is your Pokedex:")
	for pokemonName, pokemon := range p.pokemons {
		fmt.Printf("%s: %v\n", pokemonName, pokemon.BaseExperience)
	}
}
