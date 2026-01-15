package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays the next set of locations")
	fmt.Println("mapb: Displays the previous set of locations")
	return nil
}
