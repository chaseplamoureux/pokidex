package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous string
}

func main() {
	config := Config{}
	for {
		fmt.Print("pokidex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputCommand := strings.ToLower(scanner.Text())
		output, exists := commandDetail()[inputCommand] // certain types can be read directly from the function call if its the return value
		if exists {
			err := output.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("invalid command")
		}
	}
}

func commandDetail() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays locations on the pokidex map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays previous locations on the pokidex map",
			callback:    commandMapB,
		},
	}
}
