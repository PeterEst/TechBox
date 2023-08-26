package handlers

type GoHandler struct{}

func (h GoHandler) CreateProject(projectName string) {
	commandsGetter := getGoCommands
	CreateProjectCommon(projectName, "GoLang", commandsGetter)

	// ? Additional stuff can be done here
}

func getGoCommands(
	projectName string,
) []string {
	return []string{
		"touch main.go",
	}
}
