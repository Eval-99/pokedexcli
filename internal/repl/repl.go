package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Eval-99/pokedexcli/internal/commands"
	"github.com/Eval-99/pokedexcli/internal/pokecache"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := commands.Config{
		Next:    "https://pokeapi.co/api/v2/location-area",
		Cache:   *pokecache.NewCache(5 * time.Second),
		Pokedex: make(map[string]commands.PokedexEntry),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := CleanInput(scanner.Text())

		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		commandArg := ""
		if len(userInput) > 1 {
			commandArg = userInput[1]
		} else {
			commandArg = ""
		}
		command, ok := commands.GetCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.Callback(&config, commandArg)
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
