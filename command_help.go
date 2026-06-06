package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
    for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
    return nil
}