package handlers

import (
	"fmt"
	"techbox/utils"
)

type NodeHandler struct{}

func (h NodeHandler) CreateProject(projectName string) {
	fmt.Println("Creating Node project...")

	utils.MakeDir(projectName, true)

	commands := getNodeCommands(projectName)

	utils.RunCommands(commands)

	fmt.Println("Node project created!")
}

func getNodeCommands(
	projectName string,
) []string {
	return []string{
		"npm init -y",
		"touch index.js",
	}
}
