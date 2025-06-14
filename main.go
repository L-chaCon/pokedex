package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			command := scanner.Text()
			pokemons := cleanInput(command)
			fmt.Printf("Your command was: %s\n", pokemons[0])
		}
	}
}
