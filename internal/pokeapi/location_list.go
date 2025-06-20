package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locations, ok := c.pokeCache.Get(url)
	if ok {
		locationsRes := LocationAreasResponse{}
		err := json.Unmarshal(locations, &locationsRes)
		if err != nil {
			return LocationAreasResponse{}, fmt.Errorf("Not able to read cache: %w", err)
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("Error in request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("Error in response: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return LocationAreasResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("Not able to read body: %w", err)
	}

	locationsRes := LocationAreasResponse{}
	err = json.Unmarshal(body, &locationsRes)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("Not able to read body: %w", err)
	}

	c.pokeCache.Add(url, body)

	return locationsRes, nil
}
