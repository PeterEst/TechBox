package handlers

import (
	"fmt"
	"techbox/utils"
)

type NextHandler struct{}

func (h NextHandler) CreateProject(projectName string) {
	fmt.Println("Creating React project...")

	utils.MakeDir(projectName, true)

	// We can add more commands here
	commands := getNextCommands(projectName)

	utils.RunCommands(commands)

	fmt.Println("Next.JS project created!")
}

func getNextCommands(
	projectName string,
) []string {
	return []string{
		"npx create-next-app .",
		"cd " + projectName,
		"npm install",
	}
}
