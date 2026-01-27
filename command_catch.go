package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	roll := rand.Intn(100)
	difficulty := clamp(pokemon.BaseExperience, 10, 90)
	if roll > difficulty {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func clamp(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
