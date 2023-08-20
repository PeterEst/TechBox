package handlers

import (
	"fmt"
	"techbox/utils"
)

type GoHandler struct{}

func (h GoHandler) CreateProject(projectName string) {
	fmt.Println("Creating Go project...")

	utils.MakeDir(projectName, true)

	commands := getGoCommands(projectName)

	utils.RunCommands(commands)

	fmt.Println("Go project created!")
}

func getGoCommands(
	projectName string,
) []string {
	return []string{
		"touch main.go",
	}
}
