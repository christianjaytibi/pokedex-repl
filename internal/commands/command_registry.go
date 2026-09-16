package commands

import "github.com/christianjaytibi/pokedex-repl/internal/pokeapi"

type Command struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

type Config struct {
	PokeApiClient        pokeapi.Client
	CommandRegistry      map[string]Command
	CaughtPokemon        map[string]pokeapi.Pokemon
	NextLocationsURL     *string
	PreviousLocationsURL *string
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokédex.",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message.",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays names of 20 location areas in the Pokémon world.",
			Callback:    commandMapForward,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 locations.",
			Callback:    commandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Lists all Pokémon encountered within a specific location area.",
			Callback:    commandExplore,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Displays details (height, weight, stats, types) for any Pokémon you have successfully caught.",
			Callback:    commandInspect,
		},
	}
}
