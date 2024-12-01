package main

import (
	"fmt"
)

func commandMap(c *Config) error {
	locationResp, err := c.pokeapiClient.ListLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = locationResp.Next
	c.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(c *Config) error {
	if c.Previous == nil {
		return fmt.Errorf("ERROR: On first page. Cannot go further back")
	}

	locationResp, err := c.pokeapiClient.ListLocations(c.Previous)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	c.Next = locationResp.Next
	c.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
