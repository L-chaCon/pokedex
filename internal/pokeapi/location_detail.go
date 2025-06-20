package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetails(locationName string) (LocationAreaDetail, error) {
	url := baseURL + "/location-area/" + locationName
	locations, ok := c.pokeCache.Get(url)
	if ok {
		locationsRes := LocationAreaDetail{}
		err := json.Unmarshal(locations, &locationsRes)
		if err != nil {
			return LocationAreaDetail{}, fmt.Errorf("Not able to read cache: %w", err)
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaDetail{}, fmt.Errorf("Error in request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetail{}, fmt.Errorf("Error in response: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return LocationAreaDetail{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		return LocationAreaDetail{}, fmt.Errorf("Not able to read body: %w", err)
	}

	locationsRes := LocationAreaDetail{}
	err = json.Unmarshal(body, &locationsRes)
	if err != nil {
		return LocationAreaDetail{}, fmt.Errorf("Not able to read body: %w", err)
	}

	c.pokeCache.Add(url, body)

	return locationsRes, nil
}
