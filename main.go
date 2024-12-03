package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/chaseplamoureux/pokidexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

type Config struct {
	pokeapiClient	pokeapi.Client
	Next     		*string
	Previous 		*string
	Pokedex			map[string]pokeapi.Pokemon
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	config := &Config{
		pokeapiClient: pokeClient,
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
	for {
		fmt.Print("pokidex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputCommand := scanner.Text()
		formattedInputCommand := formatInputCommand(inputCommand)
		output, exists := commandDetail()[formattedInputCommand[0]] // certain types can be read directly from the function call if its the return value
		if exists {
			err := output.callback(config, formattedInputCommand)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("invalid command")
		}
	}
}

func formatInputCommand(inputCommand string) []string {
	return strings.Fields(strings.ToLower(inputCommand))
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
		"explore": {
			name:        "explore",
			description: "displays pokemon that exist in a certain location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "displays details and stats of a Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "displays Pokemon in Pokedex",
			callback:    commandPokedex,
		},
	}
}
