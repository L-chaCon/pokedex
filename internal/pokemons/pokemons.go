package pokemons

import (
	"github.com/L-chaCon/pokedex/internal/pokeapi"
)

func NewPokedex() Pokedex {
	return Pokedex{make(map[string]Pokemon)}
}

func (p *Pokedex) Add(name string, pokemon Pokemon) {
	p.pokemons[name] = pokemon
}

func ConvertFromAPIPokemon(pokemonAPI pokeapi.Pokemon) (Pokemon, error) {
	var pokemon Pokemon

	pokemon.Name = pokemonAPI.Name
	pokemon.Height = pokemonAPI.Height
	pokemon.Weight = pokemonAPI.Weight

	return pokemon, nil
}
