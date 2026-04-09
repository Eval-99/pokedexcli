package commands

import "github.com/Eval-99/pokedexcli/internal/pokecache"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous string
	Cache    pokecache.Cache
}

type (
	Locations struct {
		Next     string     `json:"next"`
		Previous string     `json:"previous"`
		Results  []Location `json:"results"`
	}
	Location struct {
		Name string `json:"name"`
	}
)

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "See next map",
			Callback:    commandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "See previous map",
			Callback:    commandMapb,
		},
	}
}
