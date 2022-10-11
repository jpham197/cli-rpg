// Package main is the entry point to the program.
package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// Account
type Account struct {
	username string
	password string
}

func (a *Account) Login() bool {
	// TODO: calls backend
	return true
}

// Base Stats
type Stats struct {
	health int64
	mana   int64
	atk    int64
	def    int64
}

// Base Creature
type Creature struct {
	Stats
	name string
}

// Main Character
type Character struct {
	Creature
}

func (c *Character) move() {
	// generates random event
	eventGenerator := rand.Intn(100)
	if eventGenerator <= 33 {
		m := Monster{}
		m.atk = 10
		m.def = 10
		m.health = 10
		m.name = "Monster"
		var buffer bytes.Buffer
		buffer.WriteString(m.name + " jumps out of a bush and attacks you!")
		// TODO: method for combat
		constructMenu("encounter")
	} else {
		var buffer bytes.Buffer
		buffer.WriteString(c.name + " moves onward")
	}
}

func fightMenu(action string) {
	switch action {
	case "F":
		fmt.Println("Fighting...")
	case "R":
		fmt.Println("Running...")
	}
}

// Monster
type Monster struct {
	Creature
}

func main() {
	acc := Account{username: os.Args[0], password: os.Args[1]}
	if !acc.Login() {
		log.Fatal("Invalid account credentials")
	}
	constructMenu("main")
}

type fn func(string)

func mainMenu(action string) {
	switch action {
	case "C":
		fmt.Println("Creating character...")
	case "VC":
		fmt.Println("Viewing character...")
	case "VI":
		fmt.Println("Viewing inventory...")
	case "E":
		fmt.Println("Exploring...")
	}
}

func constructMenu(menu string) {
	menuMap := map[string][]string{
		"main": {
			"Create character",
			"View character", // TODO: Ambiguous menu option selection
			"View inventory",
			"Explore",
		},
		"encounter": {
			"Fight",
			"Run",
		},
		"fight": {
			"Attack",
			"Defend",
			"Run",
		},
	}

	newMenuMap := map[string]fn{
		"main":  mainMenu,
		"fight": fightMenu,
	}
	// user input--------|
	//                   \/
	var input string
	fmt.Scanln(&input)
	// we would have a switch for this right?
	newMenuMap["main"](input)
	newMenuMap["fight"](input)
	// TODO: handle argument that is not in menu map -> Maybe an ENUM

	var buffer bytes.Buffer
	buffer.WriteString("Welcome to the main menu.\n")
	for _, option := range menuMap[menu] {
		buffer.WriteString(createMenuOption(option))
	}
	fmt.Println(buffer.String())
	// TODO: Handle user input for selecting an option in the menu
}

/*
Input: "View character"
Output "- [V]iew character"
*/
func createMenuOption(description string) string {
	var buffer bytes.Buffer
	str := fmt.Sprintf("- [%c]%s\n", description[0], description[1:])
	buffer.WriteString(str)
	return buffer.String()
}
