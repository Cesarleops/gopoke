package main

import (
	"errors"
	"fmt"
	"os"
)

func exitCMD(config *Config) error {
	os.Exit(1)
	return errors.New("")
}

func helpCMD(*Config) error {

	fmt.Println("Welcome to the pokedex! \n Usage: \n help: Displays a help message \n exit: Exit the Pokedex")
	return nil
}

func mapCMD(config *Config) error {

	locations, err := config.pokeClient.ListPokemons(config.Next)

	if err != nil {
		return err
	}
	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func mapbCMD(config *Config) error {

	if config.Previous == nil {
		return errors.New("you are on the first page")
	}
	locations, err := config.pokeClient.ListPokemons(config.Previous)

	if err != nil {
		return err
	}
	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
