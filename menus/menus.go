// Package menus contains all data to generate menus and handles user interaction with menus.
package menus

import (
	"bytes"
	"fmt"
)

type fn func(string)

func fightMenu(action string) {
	switch action {
	case "F":
		fmt.Println("Fighting...")
	case "R":
		fmt.Println("Running...")
	}
}

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

func ConstructMenu(menu string) {
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
