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

		cmd := exec.Command(parts[0], parts[1:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
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
