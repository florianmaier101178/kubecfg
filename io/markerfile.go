package io

import (
	"fmt"
	"os"
)

func CreateUnselectedProjectMarkerFile() int {
	if !existingPath(configDir()) {
		fmt.Printf("config directory '%s' is not existing\n", configDir())
		return 1
	}

	return createMarkerFile(fmt.Sprintf("%s/unselected", configDir()))
}

func createMarkerFile(name string) int {
	markerFile, err := os.Create(name)
	if err != nil {
		fmt.Printf("cannot create '%s'\n", name)
		return 1
	}
	defer markerFile.Close()

	return 0
}

func CreateUnselectedContextMarkerFile(projectName string) int {
	if !existingPath(configDir()) {
		fmt.Printf("config directory '%s' is not existing\n", configDir())
		return 1
	}
	if !existingPath(projectDir(projectName)) {
		fmt.Printf("project config directory '%s' is not existing\n", projectDir(projectName))
		return 1
	}

	return createMarkerFile(fmt.Sprintf("%s/unselected", projectDir(projectName)))
}
