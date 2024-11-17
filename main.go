package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {

	for {
		fmt.Print("pokidex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputCommand := scanner.Text()
		output, exists := commandDetail()[inputCommand] // certain types can be read directly from the function call if its the return value
		if exists {
			err := output.callback()
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
	}
}
