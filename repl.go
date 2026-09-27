package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(conf *config) error
}

func startRepl(conf *config) error {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		commands := conf.command

		exec, ok := commands[commandName]
		if ok {
			exec.callback(conf)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

