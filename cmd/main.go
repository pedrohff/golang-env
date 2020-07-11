package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

func main() {
	fileName := "notes.txt"
	dir, err := formatFilePathForWorkingDir(fileName)
	if err != nil {

		fmt.Println("error:", err)
		return

	}
	fmt.Println(dir)
}

func formatFilePathForWorkingDir(filename string) (string, error) {
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
