package commands

type Command struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	CommandRegistry map[string]Command
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
	}
}
