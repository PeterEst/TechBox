package main

import (
	"fmt"
	"os"
	"techbox/handlers"
)

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		handlers.HandleVersion()
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Usage: ./main <technology> <project-name>")
		return
	}

	technology := os.Args[1]
	projectName := os.Args[2]

	handler := handlers.GetHandler(technology)
	if handler == nil {
		fmt.Println("Invalid technology")
		return
	}

	handler.CreateProject(projectName)
}
