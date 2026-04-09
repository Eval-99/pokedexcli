package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func explore(cfg *Config, arg string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + arg

	var CreaturesInArea Monsters
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

		if err := json.Unmarshal(data, &CreaturesInArea); err != nil {
			return err
		}

	} else {
		if err := json.Unmarshal(info, &CreaturesInArea); err != nil {
			return err
		}
		cfg.Cache.Add(url, info)
	}

	for _, creature := range CreaturesInArea.PokemonEncounters {
		fmt.Println(creature.Pokemon.Name)
	}

	return nil
}
