package main

import "fmt"

func commandPokedex(c *Config, inputCommand []string) error {

	if len(c.Pokedex) == 0 {
		return fmt.Errorf("Your Pokedex is empty. Go catch some Pokemon!")
	}

	fmt.Println("Your Pokedex:")

	for _, v := range c.Pokedex {
		fmt.Printf("  -%s\n", v.Name)
	}

	return nil
}