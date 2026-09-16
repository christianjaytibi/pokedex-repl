package commands

import (
	"fmt"
)

func commandMapForward(cfg *Config, args ...string) error {
	url := cfg.NextLocationsURL
	locationAreas, err := cfg.PokeApiClient.ListLocAreas(url)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationAreas.Next
	cfg.PreviousLocationsURL = locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapBack(cfg *Config, args ...string) error {
	if cfg.PreviousLocationsURL == nil {
		return fmt.Errorf("Cannot go back any further.")
	}

	url := cfg.PreviousLocationsURL
	locationAreas, err := cfg.PokeApiClient.ListLocAreas(url)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationAreas.Next
	cfg.PreviousLocationsURL = locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
