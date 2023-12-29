package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func commandHelp() error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex

`)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func printPrompt() {
	fmt.Print("Pokedex > ")
}

func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

func main() {
	commands := map[string]cliCommand{
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
	}

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		input := cleanInput(reader.Text())
		command, ok := commands[input]
		if !ok {
			printUnknown(input)
			printPrompt()
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err.Error())
		}
		printPrompt()
	}

}
