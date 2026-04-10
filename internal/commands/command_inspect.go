package commands

import "fmt"

func inspect(cfg *Config, arg string) error {
	creature, ok := cfg.Pokedex[arg]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	stats := map[string]int{}
	for _, stat := range creature.Stats {
		name := stat.Name.StatNam
		val := stat.Stat
		stats[name] = val
	}

	fmt.Printf("Name: %s\n", creature.Name)
	fmt.Printf("Height: %d\n", creature.Height)
	fmt.Printf("Weight: %d\n", creature.Weight)

	fmt.Println("Stats:")
	fmt.Printf("  -hp: %d\n", stats["hp"])
	fmt.Printf("  -attack: %d\n", stats["attack"])
	fmt.Printf("  -defense: %d\n", stats["defense"])
	fmt.Printf("  -special-attack: %d\n", stats["special-attack"])
	fmt.Printf("  -special-defense: %d\n", stats["special-defense"])
	fmt.Printf("  -speed: %d\n", stats["speed"])

	fmt.Println("Types:")
	for i := range len(creature.Types) {
		fmt.Printf("  - %s\n", creature.Types[i].Type.Name)
	}

	return nil
}
