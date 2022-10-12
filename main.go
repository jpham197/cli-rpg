// Package main is the entry point to the program.
package main

import (
	account "clirpg/main/accounts"
	menu "clirpg/main/menus"
	"log"
	"os"
)

func main() {
	acc := account.Account{Username: os.Args[0], Password: os.Args[1]}
	if !acc.Login() {
		log.Fatal("Invalid account credentials")
	}

	// TODO: This function will freeze the CLI
	menu.ConstructMenu("main")
}
