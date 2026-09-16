package main

import (
	"time"

	"github.com/christianjaytibi/pokedex-repl/internal/commands"
	"github.com/christianjaytibi/pokedex-repl/internal/pokeapi"
	"github.com/christianjaytibi/pokedex-repl/internal/repl"
)

func main() {
	clientTimeout := time.Second * 5
	cacheInterval := time.Minute * 5
	pokeClient := pokeapi.NewClient(clientTimeout, cacheInterval)

	cfg := &commands.Config{
		CommandRegistry: commands.GetCommands(),
		CaughtPokemon:   make(map[string]pokeapi.Pokemon),
		PokeApiClient:   pokeClient,
	}

	repl.Start(cfg)
}
