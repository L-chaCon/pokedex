package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	pokemons := strings.Fields(lowerText)
	return pokemons
}
