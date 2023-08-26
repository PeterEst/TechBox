package handlers

type NodeHandler struct{}

func (h NodeHandler) CreateProject(projectName string) {
	commandsGetter := getNodeCommands
	CreateProjectCommon(projectName, "Node", commandsGetter)

	// ? Additional stuff can be done here
}

func getNodeCommands(
	projectName string,
) []string {
	return []string{
		"npm init -y",
		"touch index.js",
	}
}
