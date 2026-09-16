package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocAreas(pageUrl *string) (RespShallowLocAreas, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if data, ok := c.cache.Get(url); ok {
		var locationAreas RespShallowLocAreas
		if err := json.Unmarshal(data, &locationAreas); err != nil {
			return RespShallowLocAreas{}, err
		}

		return locationAreas, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return RespShallowLocAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocAreas{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocAreas{}, err
	}
	c.cache.Add(url, data)

	var locationAreas RespShallowLocAreas
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return RespShallowLocAreas{}, err
	}

	return locationAreas, nil
}
