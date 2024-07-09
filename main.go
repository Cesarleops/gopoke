package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cesarleops/pockedex/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

type Config struct {
	pokeClient pokeapi.Client
	Previous   *string
	Next       *string
}

func cleanInput(text string) []string {

	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

func checkCommand(options map[string]cliCommand, comm string) (cliCommand, error) {

	v, ok := options[comm]

	if !ok {
		return cliCommand{}, errors.New("not found")
	}
	return v, nil
}

func startREPL(config *Config) {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCMD,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCMD,
		},
		"map": {
			name:        "map",
			description: "See 20 pokemon locations",
			callback:    mapCMD,
		},
		"mapb": {
			name:        "mapb",
			description: "See previoues 20 pokemon locations",
			callback:    mapbCMD,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("pokedex > ")
	for scanner.Scan() {
		line := scanner.Text()
		cleanedLine := cleanInput(line)
		cmd, err := checkCommand(commands, cleanedLine[0])
		if err != nil {
			fmt.Println(err)
			break
		}

		cmd.callback(config)
	}

}

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)

	config := &Config{
		pokeClient: pokeClient,
	}
	startREPL(config)
}
