package main

import (
	"github.com/christianjaytibi/pokedex-repl/internal/commands"
	"github.com/christianjaytibi/pokedex-repl/internal/repl"
)

func main() {
	cfg := &commands.Config{
		CommandRegistry: commands.GetCommands(),
	}

	repl.Start(cfg)
}
