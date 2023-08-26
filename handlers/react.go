package handlers

type ReactHandler struct{}

func (h ReactHandler) CreateProject(projectName string) {
	commandsGetter := getReactCommands
	CreateProjectCommon(projectName, "ReactJS", commandsGetter)

	// ? Additional stuff can be done here
}

func getReactCommands(
	projectName string,
) []string {
	return []string{
		"npx create-react-app .",
	}
}
