package main

type config struct {
	command		map[string]cliCommand 	
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:		"exit",
			description:	"Exit the Pokedex",
			callback:	commandExit,
		},
		"help": {
			name:		"help",
			description:	"Displays a help message",
			callback:	commandHelp,
		},
	}
}

func main() {
	conf := &config{
		command: getCommands(),
	}

	startRepl(conf)
}
