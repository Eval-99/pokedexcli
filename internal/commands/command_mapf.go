package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapf(cfg *Config, arg string) error {
	if cfg.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	var locations Locations
	if info, ok := cfg.Cache.Get(cfg.Next); !ok {
		res, err := http.Get(cfg.Next)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.Cache.Add(cfg.Next, data)

		if err := json.Unmarshal(data, &locations); err != nil {
			return err
		}

	} else {
		if err := json.Unmarshal(info, &locations); err != nil {
			return err
		}
		cfg.Cache.Add(cfg.Next, info)
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}
