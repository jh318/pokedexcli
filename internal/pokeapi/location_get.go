package pokeapi

import (
    "encoding/json"
    "io"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	if data, ok := c.cache.Get(url); ok {
		var location Location
		if err := json.Unmarshal(data, &location); err != nil {
			return Location{}, err
		}
		return location, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	var location Location
	if err := json.Unmarshal(body, &location); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, body)
	return location, nil
}