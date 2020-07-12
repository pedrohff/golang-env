package pkg

import (
	"errors"
	"os"
	"strings"
)

// FormatFilePathForWorkingDir Formats a path for a given file, appending it to the current directory
func FormatFilePathForWorkingDir(filename string) (string, error) {
	fileHasExtension := strings.Contains(filename, ".")
	if !fileHasExtension {
		return "", errors.New("please declare a file extension")
	}

	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return workingDirectory + filename, nil

}
