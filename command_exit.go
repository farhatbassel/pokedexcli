package main

import "os"

func commandExit(config *config, arg string) error {
    os.Exit(0)
    return nil
}
