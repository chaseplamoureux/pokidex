package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocationDetails(id string) (locationDetails, error) {
	url := baseURL + "/location-area/" + id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationDetails{}, err
	}

	if val, ok := c.cache.Get(url); ok {
		locationDetailsResp := locationDetails{}
		err := json.Unmarshal(val, &locationDetailsResp)
		if err != nil {
			return locationDetails{}, err
		}
		return locationDetailsResp, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationDetails{}, err
	}

	locationDetailsResp := locationDetails{}
	err = json.Unmarshal(dat, &locationDetailsResp)
	if err != nil {
		return locationDetails{}, err
	}

	return locationDetailsResp, nil
}
