package pokeapi

import (
	"encoding/json"
	"io"
)

//import "net/http"

func (c *Client) GetLocationAreas(url string) (RespShallowLocations, error) {
	if data, ok := c.cache.Get(url); ok {
		var locations RespShallowLocations
		if err := json.Unmarshal(data, &locations); err != nil {
			return RespShallowLocations{}, err
		}
		return locations, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, body)

	var locations RespShallowLocations
	if err := json.Unmarshal(body, &locations); err != nil {
		return RespShallowLocations{}, err
	}

	return locations, nil
}
