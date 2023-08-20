package handlers

type ProjectHandler interface {
	CreateProject(projectName string)
}

/**
 * GetHandler returns the handler for the given technology
 */
func GetHandler(technology string) ProjectHandler {
	switch technology {
	case "react":
		return &ReactHandler{}
	case "next":
		return &NextHandler{}
	default:
		return nil
	}
}
