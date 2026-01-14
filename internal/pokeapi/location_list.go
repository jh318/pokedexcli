package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

//import "net/http"

func GetLocationAreas(url string) (RespShallowLocations, error) {
	// TODO: make HTTP request, decode JSON
	res, err := http.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	var locations RespShallowLocations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locations, nil
}
