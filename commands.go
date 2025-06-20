package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/L-chaCon/pokedex/internal/config"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config, ...string) error
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
		"explore": {
			name:        "explore <location name>",
			description: "Shows list of pokemon in a area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Print the poxedex that you have",
			callback:    commandPrintDex,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect the information of a pokemon",
			callback:    commandInspect,
		},
	}

}

func commandExit(cfg *config.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config.Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMapF(cfg *config.Config, args ...string) error {
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

func commandMapB(cfg *config.Config, args ...string) error {
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

func commandExplore(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}
	location := args[0]

	fmt.Printf("Exploring %s...\n", location)

	locationDetails, err := cfg.PokeapiClient.GetLocationDetails(location)
	if err != nil {
		return fmt.Errorf("Error in GetLocationDetails: %w", err)
	}

	for _, encounter := range locationDetails.PokemonEncounters {
		fmt.Printf("\t- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon to catch")
	}
	pokemon := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	pokemonInfo, err := cfg.PokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("Error in GetPokemon: %w", err)
	}

	if pokemonInfo.BaseExperience < rand.IntN(cfg.RandomRoof) {
		cfg.Pokedex.Add(pokemonInfo.Name, pokemonInfo)
		fmt.Println("Captured ;)")
	} else {
		fmt.Println("BAD LUCK!")
	}

	return nil
}

func commandPrintDex(cfg *config.Config, args ...string) error {
	cfg.Pokedex.PrintPokemons()
	return nil
}

func commandInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon to inspect")
	}
	pokemon := args[0]

	value, ok := cfg.Pokedex.Get(pokemon)
	if !ok {
		return errors.New("You have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", value.Name)
	fmt.Printf("Height: %d\n", value.Height)
	fmt.Printf("Weight: %d\n", value.Weight)
	fmt.Println("Stats:")
	for _, value := range value.Stats {
		fmt.Printf(" - %s: %d\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Println("Types:")
	for _, value := range value.Types {
		fmt.Printf(" - %s\n", value.Type.Name)
	}
	return nil
}
