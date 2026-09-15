package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocAreas(pageUrl *string) (RespShallowLocAreas, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
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

	var locationAreas RespShallowLocAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationAreas); err != nil {
		return RespShallowLocAreas{}, err
	}

	return locationAreas, nil
}
