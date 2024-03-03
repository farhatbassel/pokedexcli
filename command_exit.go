package main

import "os"

func commandExit(config *config, area string) error {
    os.Exit(0)
    return nil
}
