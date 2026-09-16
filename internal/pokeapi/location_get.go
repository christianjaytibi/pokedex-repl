package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (LocationArea, error) {
	url := baseURL + "/location-area/" + location

	if data, ok := c.cache.Get(url); ok {
		locArea := LocationArea{}
		if err := json.Unmarshal(data, &locArea); err != nil {
			return LocationArea{}, err
		}

		return locArea, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return LocationArea{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, data)

	var locArea LocationArea
	if err := json.Unmarshal(data, &locArea); err != nil {
		return LocationArea{}, err
	}

	return locArea, nil
}
