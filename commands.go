package main

import (
	"fmt"
	"github.com/L-chaCon/pokedex/internal/config"
	"github.com/L-chaCon/pokedex/internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays page of map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back a page of map locations",
			callback:    commandMapB,
		},
	}
}

func commandExit(c *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *config.Config) error {
	err := pokeapi.NextLocationAreaPage(c)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func commandMapB(c *config.Config) error {
	err := pokeapi.PreviousLocationAreasPage(c)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
