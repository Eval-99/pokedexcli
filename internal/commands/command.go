package commands

import (
	"github.com/Eval-99/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
}

type Config struct {
	Next     string
	Previous string
	Cache    pokecache.Cache
	Pokedex  map[string]PokedexEntry
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

type (
	Monsters struct {
		PokemonEncounters []monster `json:"pokemon_encounters"`
	}
	monster struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	}
)

type (
	PokedexEntry struct {
		BaseExp int    `json:"base_experience"`
		Name    string `json:"name"`
		Height  int    `json:"height"`
		Weight  int    `json:"weight"`
		Types   []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
		Stats []struct {
			Name struct {
				StatNam string `json:"name"`
			} `json:"stat"`
			Stat int `json:"base_stat"`
		} `json:"stats"`
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
		"explore": {
			Name:        "explore",
			Description: "See monsters in specified area",
			Callback:    explore,
		},
		"catch": {
			Name:        "catch",
			Description: "See monsters in specified area",
			Callback:    catch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect monsters in Pokedex",
			Callback:    inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show all monsters in Pokedex",
			Callback:    pokedex,
		},
	}
}
