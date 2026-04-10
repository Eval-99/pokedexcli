package commands

import "fmt"

func pokedex(cfg *Config, arg string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty...")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for poke := range cfg.Pokedex {
		fmt.Printf(" - %s", poke)
	}

	return nil
}
