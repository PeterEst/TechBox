package handlers

import (
	"fmt"
	"techbox/utils"
)

func CreateProjectCommon(projectName, successMessage string, commandsGetter func(string) []string) {
	fmt.Printf("Creating %s project...\n", successMessage)

	utils.MakeDir(projectName, true)

	// Get technology-specific commands
	commands := commandsGetter(projectName)

	// Run technology-specific commands
	utils.RunCommands(commands)

	fmt.Printf("%s project created!\n", successMessage)
}
