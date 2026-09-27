package main

import (
	"fmt"
)

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, elem := range conf.command {
		fmt.Printf("%s: %s\n", key, elem.description)
	}
	return nil
}

