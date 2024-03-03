package main

import (
    "github.com/fatih/color"
    "fmt"
)

func commandHelp (config *config) error {
    commands := getCliCommands()

	helpMessage := "Welcome to the Pokedex!\nUsage:\n\n"

	for cmd, cliCommand := range commands {
		coloredCmd := getColorString(cmd)
		helpMessage += fmt.Sprintf("%s: %s\n", coloredCmd, cliCommand.description)
	}

	fmt.Println(helpMessage)
	return nil
}

// getColorString returns a colored string for a command
func getColorString(cmd string) string {
	switch cmd {
	case "help":
		return color.GreenString(cmd)
	case "exit":
		return color.RedString(cmd)
	case "map":
		return color.BlueString(cmd)
	case "mapb":
		return color.YellowString(cmd)
	default:
		return cmd
	}
}
