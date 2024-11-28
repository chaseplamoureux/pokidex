package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type mapResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// redo commandMap and commandMapb
// Take the actual http call and json unmarshal and result printing into itw own function.
// return Previous and Next values from that function to commandMap and commandMapB and update config struct.

func commandMap(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Next != "" {
		url = c.Next
	}

	err := printLocations(c, url)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func commandMapB(c *Config) error {
	if c.Previous == "" {
		return fmt.Errorf("ERROR: On first page. Cannot go further back")
	}

	err := printLocations(c, c.Previous)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func printLocations(c *Config, url string) error {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	MappedData := mapResponse{}
	err = json.Unmarshal(body, &MappedData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	c.Next = MappedData.Next
	c.Previous = MappedData.Previous

	for _, item := range MappedData.Results {
		fmt.Println(item.Name)
	}
	return nil
}
