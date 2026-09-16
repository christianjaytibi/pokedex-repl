package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if data, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
