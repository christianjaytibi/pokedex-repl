package commands

import "fmt"

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Println("Your Pokédex:")
	for pokemon := range cfg.CaughtPokemon {
		fmt.Printf("\t- %s\n", pokemon)
	}

	return nil
}
