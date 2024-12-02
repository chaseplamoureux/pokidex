package main

import (
	"fmt"
)

func commandHelp(c *Config, inputCommand []string) error {
	fmt.Println("Usage: ")
	fmt.Println("")
	commands := commandDetail()

	for k, v := range commands {
		fmt.Println(fmt.Sprintf("%s: %s", k, v.description))
	}
	return nil
}
