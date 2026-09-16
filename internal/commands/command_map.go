package commands

import (
	"fmt"
)

func commandMap(cfg *Config, args ...string) error {
	url := cfg.NextLocationsURL
	if cfg.CurrentCommand == "mapb" {
		if cfg.PreviousLocationsURL == nil {
			return fmt.Errorf("Cannot go back any further.")
		}
		url = cfg.PreviousLocationsURL
	}

	locationAreas, err := cfg.PokeApiClient.ListLocAreas(url)
	if err != nil {
		return err
	}

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	cfg.NextLocationsURL = locationAreas.Next
	cfg.PreviousLocationsURL = locationAreas.Previous

	return nil
}
