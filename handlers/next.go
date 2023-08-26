package handlers

type NextHandler struct{}

func (h NextHandler) CreateProject(projectName string) {
	commandsGetter := getNextCommands
	CreateProjectCommon(projectName, "NextJS", commandsGetter)

	// ? Additional stuff can be done here
}

func getNextCommands(
	projectName string,
) []string {
	return []string{
		"npx create-next-app@latest .",
	}
}
