package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*urlData) error
}

type urlData struct {
	next     string
	previous string
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

func commandExit(urlData *urlData) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(urlData *urlData) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(urls *urlData) error {
	if urls.next == "" {
		fmt.Println("you're on the last page")
		return nil
	}
	res, err := http.Get(urls.next)
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

	urls.next = locations.Next
	urls.previous = locations.Previous

	return nil
}

func commandMapb(urls *urlData) error {
	if urls.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(urls.previous)
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

	urls.next = locations.Next
	urls.previous = locations.Previous

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "See next map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "See previous map",
			callback:    commandMapb,
		},
	}
}
