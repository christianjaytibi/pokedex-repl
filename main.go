package main

import (
	"time"

	"github.com/christianjaytibi/pokedex-repl/internal/commands"
	"github.com/christianjaytibi/pokedex-repl/internal/pokeapi"
	"github.com/christianjaytibi/pokedex-repl/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	cfg := &commands.Config{
		CommandRegistry: commands.GetCommands(),
		PokeApiClient:   pokeClient,
	}

	repl.Start(cfg)
}
