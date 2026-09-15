package commands

import "github.com/christianjaytibi/pokedex-repl/internal/pokeapi"

type Command struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	PokeApiClient        pokeapi.Client
	CommandRegistry      map[string]Command
	CurrentCommand       string
	NextLocationsURL     *string
	PreviousLocationsURL *string
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex.",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message.",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays names of 20 location areas in the Pokemon world.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 locations.",
			Callback:    commandMap,
		},
	}
}
