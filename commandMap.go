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

func commandMap(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Next != "" {
		url = c.Next
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
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
	}
	c.Next = MappedData.Next
	if MappedData.Previous != "" {
		c.Previous = MappedData.Previous
	}

	for _, item := range MappedData.Results {
		fmt.Println(item.Name)
	}
	return nil
}

func commandMapB(c *Config) error {
	if c.Previous == "" {
		return fmt.Errorf("ERROR: On first page. Cannot go further back")
	}
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Previous != "" {
		url = c.Previous
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
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
	}
	c.Previous = MappedData.Previous
	if MappedData.Next != "" {
		c.Next = MappedData.Next
	}

	for _, item := range MappedData.Results {
		fmt.Println(item.Name)
	}
	return nil
}
