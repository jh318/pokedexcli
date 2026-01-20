package pokeapi

func (c *Client) GetLocation(locationName string) (Location, error) {
	if data, ok : c.cache.Get(locationName); ok {
		var location Location
		if err := json.Unmarshal(data, &locations); err != nil {
			return Location{}, err
		}
		return locations, nil
	}
}