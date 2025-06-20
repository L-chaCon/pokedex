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

func (p *Pokedex) Get(name string) (pokeapi.Pokemon, bool) {
	value, ok := p.pokemons[name]
	return value, ok
}

func (p *Pokedex) PrintPokemons() {
	fmt.Println("This is your Pokedex:")
	for pokemonName, _ := range p.pokemons {
		fmt.Printf(" - %s\n", pokemonName)
	}
}
