package main

import (
	"bufio"
	"farhatbassel/pokedexcli/internal/pokeapi"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type config struct {
    pokeapiClient   pokeapi.Client
    nextLocationURL *string
    prevLocationURL *string
}

func repl(config *config) {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Pokedex> ")
        input, err := reader.ReadString('\n')
        if err != nil {
            error := fmt.Sprintf("Error occured when read the string value: %v", err)
            fmt.Println(error)
            break
        }

        input = input[:len(input) - 1] // Remove the newline char from input
        inputCommand := cleanInput(input)[0]

        if command, ok := getCliCommands()[inputCommand]; !ok {
            fmt.Println("Unknown command. Please use " + color.GreenString("help"))
        } else { 
            if err := command.callback(config); err != nil {
                fmt.Printf("Error when executing command %v\n %v\n", input, err)
            }
        }
    }
}

func cleanInput(input string) []string {
    output := strings.ToLower(input)
    words := strings.Fields(output)
    return words
}

type cliCommand struct {
    name        string
    description string
    callback    func(*config) error
}

func getCliCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name: "help",
            description: "Display the Pokedex commands",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
        "map": {
            name: "map",
            description: "Display 20 locations of the Pokemon world",
            callback: commandMap,
        },
        "mapb": {
            name: "map back",
            description: "Display the previous 20 locations of the Pokemon world",
            callback: commandMapBack,
        },
    }
}

