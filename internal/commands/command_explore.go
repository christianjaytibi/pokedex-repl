package commands

import "fmt"

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("You must provide a location area.")
	}
	locationName := args[0]

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon: ")

	location, err := cfg.PokeApiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
