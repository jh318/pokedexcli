package main

import (
	"fmt"

	"github.com/jh318/pokedexcli/internal/pokeapi"
)

//import "net/http"

func commandMap(cfg *config) error {
	const locationAreaURL = "https://pokeapi.co/api/v2/location-area?limit=20"

	var url string
	if cfg.nextURL == nil {
		url = locationAreaURL
	} else {
		url = *cfg.nextURL
	}

	resp, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)

	}

	cfg.nextURL = resp.Next
	cfg.previousURL = resp.Previous

	return nil
}

func commandMapb(cfg *config) error {
	var url string
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = *cfg.previousURL
	}

	resp, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)

	}

	cfg.nextURL = resp.Next
	cfg.previousURL = resp.Previous

	return nil
}
