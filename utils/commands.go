package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/**
 * Run a list of commands.
 * Example: ["npx create-react-app .", "npm install"] -> npx create-react-app . && npm install
 */
func RunCommands(commands []string) {
	for _, cmd := range commands {
		parts := splitCommand(cmd)
		out, err := exec.Command(parts[0], parts[1:]...).CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Output:", string(out))
			os.Exit(1)
		}

		// TODO: We can show a progress bar here
	}
}

/**
 * Split a command into its parts.
 * Example: "npx create-react-app ." -> ["npx", "create-react-app", "."]
 *
 * Reason for this function: exec.Command() expects the command and its arguments
 */
func splitCommand(command string) []string {
	parts := []string{}
	for _, part := range strings.Split(command, " ") {
		if part != "" {
			parts = append(parts, part)
		}
	}
	return parts
}
