package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapf(cfg *Config) error {
	if cfg.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	res, err := http.Get(cfg.Next)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locations Locations
	if err := json.Unmarshal(data, &locations); err != nil {
		return err
	}

	defer res.Body.Close()

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}
