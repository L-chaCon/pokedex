package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/L-chaCon/pokedex/internal/config"
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
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back a page of map locations",
			callback:    commandMapB,
		},
	}
}

func commandExit(cfg *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMapF(cfg *config.Config) error {
	locationList, err := cfg.PokeapiClient.GetLocationAreas(cfg.NextLocationsURL)
	if err != nil {
		return fmt.Errorf("Error running Mapf: %w", err)
	}
	cfg.NextLocationsURL = locationList.Next
	cfg.PrevLocationsURL = locationList.Previous

	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *config.Config) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("You're on the first page")
	}

	locationList, err := cfg.PokeapiClient.GetLocationAreas(cfg.PrevLocationsURL)
	if err != nil {
		return fmt.Errorf("Error running Mapb: %w", err)
	}
	cfg.NextLocationsURL = locationList.Next
	cfg.PrevLocationsURL = locationList.Previous

	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}

	return nil
}
