package io

import (
	"fmt"
	"os"
)

func CreateConfigProjectSymlink() int {
	if !existingPath(configDir()) {
		fmt.Printf("project config directory '%s' is not existing\n", configDir())
		return 1
	}

	return createSymlink(unselectedProjectFile(), configSymlink())
}

func createSymlink(target string, symlink string) int {
	err := os.Symlink(target, symlink)
	if err != nil {
		fmt.Printf("cannot create '%s'\n", symlink)
		return 1
	}

	return 0
}

func CreateProjectContextSymlink(projectName string) int {
	if !existingPath(projectDir(projectName)) {
		fmt.Printf("project config directory '%s' is not existing\n", projectDir(projectName))
		return 1
	}

	return createSymlink(unselectedContextFile(projectName), projectSymlink(projectName))
}
