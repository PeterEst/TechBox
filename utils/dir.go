package utils

import "os"

/**
 * Create a directory if it doesn't exist and enter it.
 */
func MakeDir(dirName string, enterDir bool) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.Mkdir(dirName, 0755)
	}

	if enterDir {
		os.Chdir(dirName)
	}
}
