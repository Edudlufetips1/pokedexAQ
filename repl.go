package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand {
		"exit": {
			name:        	"exit",
			description: 	"Exit the Pokedex",
			callback:    	commandExit,
		},
		"help": {
			name:			"help",
			description: 	"Displays a help message",
			callback:		commandHelp,
			},
	}
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		if command, exists := commands[commandName]; exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
		fmt.Println("Unknown command")}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}
