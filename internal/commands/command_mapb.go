package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var locations Locations
	if info, ok := cfg.Cache.Get(cfg.Previous); !ok {
		res, err := http.Get(cfg.Previous)
		if err != nil {
			return err
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cfg.Cache.Add(cfg.Previous, data)

		if err := json.Unmarshal(data, &locations); err != nil {
			return err
		}

		defer res.Body.Close()
	} else {
		if err := json.Unmarshal(info, &locations); err != nil {
			return err
		}
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}
