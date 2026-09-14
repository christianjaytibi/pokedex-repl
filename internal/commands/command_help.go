package commands

import "fmt"

func commandHelp(cfg *Config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, c := range cfg.CommandRegistry {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}
	fmt.Println()
	return nil
}
