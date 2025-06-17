package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	Next    string
	Preview string
}

var BaseConfig config

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandWord := words[0]
		command, exist := getCommands()[commandWord]
		if exist {
			err := command.callback(&BaseConfig)
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
