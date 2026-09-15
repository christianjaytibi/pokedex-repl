package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/christianjaytibi/pokedex-repl/internal/commands"
)

func Start(cfg *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}

		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if cmd, exists := cfg.CommandRegistry[commandName]; exists {
			cfg.CurrentCommand = commandName
			err := cmd.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	return strings.Fields(lowered)
}
