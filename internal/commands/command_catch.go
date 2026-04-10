package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func catch(cfg *Config, arg string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + arg

	var Creature PokedexEntry
	if info, ok := cfg.Cache.Get(url); !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.StatusCode > 299 {
			fmt.Println("Does not exist")
			return nil
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.Cache.Add(url, data)

		if err := json.Unmarshal(data, &Creature); err != nil {
			return err
		}

	} else {
		if err := json.Unmarshal(info, &Creature); err != nil {
			return err
		}
		cfg.Cache.Add(url, info)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", Creature.Name)
	chance := rand.Intn(150)
	if chance >= Creature.BaseExp {
		fmt.Printf("%s was caught!\n", Creature.Name)
		cfg.Pokedex[Creature.Name] = Creature
	} else {
		fmt.Printf("%s escaped!\n", Creature.Name)
	}

	return nil
}
