package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	locations, ok := c.pokeCache.Get(url)
	if ok {
		pokemonRes := Pokemon{}
		err := json.Unmarshal(locations, &pokemonRes)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Not able to read cache: %w", err)
		}
		return pokemonRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error in request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error in response: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		return Pokemon{}, fmt.Errorf("Not able to read body: %w", err)
	}

	pokemonRes := Pokemon{}
	err = json.Unmarshal(body, &pokemonRes)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Not able to read body: %w", err)
	}

	c.pokeCache.Add(url, body)

	return pokemonRes, nil
}
