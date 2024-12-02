package main

import (
	"os"
)

func commandExit(c *Config, inputCommand []string) error {
	os.Exit(0)
	return nil
}
