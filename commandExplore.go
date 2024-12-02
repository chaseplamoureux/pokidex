package main

import "fmt"

func commandExplore(c *Config, inputCommand []string) error {

	if len(inputCommand) < 2 {
		return fmt.Errorf("explore command requires a location name or id\nUsage: explore mt-coronet-5f") 
	}
	id := inputCommand[1]
	locDetails, err := c.pokeapiClient.ListLocationDetails(id)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + id)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locDetails.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}
