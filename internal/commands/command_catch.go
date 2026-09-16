package commands

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide the name of the Pokémon to catch.")
	}

	pokemonName := args[0]
	if _, ok := cfg.CaughtPokemon[pokemonName]; ok {
		return errors.New("Pokémon is already in the Pokédex. Run \"inspect <pokemon-name>\" for details.")
	}

	pokemon, err := cfg.PokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	base_formula := float64(105) - float64(pokemon.BaseExperience)/float64(3)
	catch_probability := math.Max(5, math.Min(95, base_formula)) / float64(100)

	r := rand.Float64()
	if r <= catch_probability {
		fmt.Printf("Gotcha! %s was caught!\n", pokemon.Name)
		cfg.CaughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("Oh no! The wild %s escaped!\n", pokemon.Name)
	}

	return nil
}
