package main

import "fmt"

func commandInspect(c *Config, inputCommand []string) error {
	pokemonName := inputCommand[1]

	pokemon, ok := c.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught %s yet", pokemonName)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")

	for _, v := range pokemon.Stats {
		baseStat := v.BaseStat
		statName := v.Stat.Name
		fmt.Printf("  -%s: %d\n", statName, baseStat)
	}

	fmt.Println("Types:")

	for _, v := range pokemon.Types {
		typeName := v.Type.Name
		fmt.Printf("  - %s\n", typeName)
	}
	return nil
}
