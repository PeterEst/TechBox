package handlers

import (
	"fmt"
	"techbox/utils"
)

type ReactHandler struct{}

func (h ReactHandler) CreateProject(projectName string) {
	fmt.Println("Creating React project...")

	utils.MakeDir(projectName, true)

	// We can add more commands here
	commands := getReactCommands(projectName)

	utils.RunCommands(commands)

	fmt.Println("React project created!")
}

func getReactCommands(
	projectName string,
) []string {
	return []string{
		"npx create-react-app .",
	}
}
