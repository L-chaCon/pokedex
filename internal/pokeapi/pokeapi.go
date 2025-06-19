package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/L-chaCon/pokedex/internal/config"
	"io"
	"net/http"
)

var LOCATION_AREA_URL string = "https://pokeapi.co/api/v2/location-area/"

type locationAreasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationAreas(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Not able to get location area: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, fmt.Errorf("Not able to read body: %w", err)
	}

	locationsRes := locationAreasResponse{}
	err = json.Unmarshal(body, &locationsRes)
	if err != nil {
		return nil, fmt.Errorf("Not able to read body: %w", err)
	}
	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}
	pre := ""
	if locationsRes.Previous != nil {
		pre = fmt.Sprintf("%v", locationsRes.Previous)
	}

	return []string{locationsRes.Next, pre}, nil
}

func PreviousLocationAreasPage(c *config.Config) error {
	if c.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	url := fmt.Sprintf("%v", c.Previous)

	next_list, err := getLocationAreas(url)
	if err != nil {
		return fmt.Errorf("Error in previews: %w", err)
	}
	c.Next = next_list[0]
	c.Previous = next_list[1]

	return nil
}

func NextLocationAreaPage(c *config.Config) error {
	if c.Next == "" {
		c.Next = LOCATION_AREA_URL
	}

	next_list, err := getLocationAreas(c.Next)
	if err != nil {
		return fmt.Errorf("Error in next: %w", err)
	}
	c.Next = next_list[0]
	c.Previous = next_list[1]

	return nil
}
