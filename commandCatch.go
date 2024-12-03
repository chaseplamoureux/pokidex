package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *Config, inputCommand []string) error {
	if len(inputCommand) < 2 {
		return fmt.Errorf("catch command requires a Pokemon name or id\nUsage: catch pikachu")
	}
	id := inputCommand[1]

	_, ok := c.Pokedex[id]
	if ok {
		return fmt.Errorf("You have already caught this Pokemon!")
	}
	pokemonDetails, err := c.pokeapiClient.GetPokemonDetails(id)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a pokeball at %s...\n", pokemonDetails.Name)
	const threshold = 50
	randNum := rand.Intn(pokemonDetails.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("%s escaped!", pokemonDetails.Name)
	}
	fmt.Printf("%s was caught!\n", pokemonDetails.Name)
	c.Pokedex[id] = pokemonDetails
	return nil
}
