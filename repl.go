package main

import (
	"bufio"
	"fmt"
	"github.com/L-chaCon/pokedex/internal/config"
	"os"
	"strings"
)

func startRepl(cfg *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		var parameters string = ""
		if len(words) > 1 {
			parameters = words[1]
		}

		commandWord := words[0]
		command, exist := getCommands()[commandWord]
		if exist {
			err := command.callback(cfg, parameters)
			if err != nil {
				fmt.Printf("Error in callback: %v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	pokemons := strings.Fields(lowerText)
	return pokemons
}
