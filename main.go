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
		command := scanner.Text()
		fmt.Println(command)
	}
}
