package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if data, ok := c.cache.Get(url); ok {
		var pokemon Pokemon
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	return pokemon, nil
}
