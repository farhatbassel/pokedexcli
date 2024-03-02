package main

import (
    "github.com/fatih/color"
    "fmt"
)

func commandHelp () error {
    helpMessage := `
Welcome to the Pokedex!
Usage:

` + color.GreenString("help") + `: Displays a help message
` + color.RedString("exit") + `: Exit the Pokedex
`

    fmt.Println(helpMessage)
    return nil
}
