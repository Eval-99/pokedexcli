package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := urlData{next: "https://pokeapi.co/api/v2/location-area"}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())

		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

	}
}

func cleanInput(text string) []string {
	stringsSlc := []string{}

	for string := range strings.FieldsSeq(text) {
		stringsSlc = append(stringsSlc, strings.ToLower(string))
	}

	return stringsSlc
}
