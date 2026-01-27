package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays the next set of locations")
	fmt.Println("mapb: Displays the previous set of locations")
	fmt.Println("explore: List the Pokemon in a location area")
	fmt.Println("catch: Catch a Pokemon")
	fmt.Println("inspect: inspect a captured Pokemon")
	fmt.Println("pokedex: list captured Pokemon")
	return nil
}
