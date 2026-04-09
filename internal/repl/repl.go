package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Eval-99/pokedexcli/internal/commands"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := commands.Config{Next: "https://pokeapi.co/api/v2/location-area"}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := CleanInput(scanner.Text())

		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		command, ok := commands.GetCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.Callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

	}
}

func CleanInput(text string) []string {
	stringsSlc := []string{}

	for string := range strings.FieldsSeq(text) {
		stringsSlc = append(stringsSlc, strings.ToLower(string))
	}

	return stringsSlc
}
