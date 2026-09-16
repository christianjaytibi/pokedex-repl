package commands

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide the name of the Pokémon to inspect.")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.CaughtPokemon[pokemonName]
	if !ok {
		return errors.New("You have not caught that Pokémon yet.")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("\t- %s\n", pokeType.Type.Name)
	}

	return nil
}
