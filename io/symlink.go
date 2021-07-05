package io

import (
	"fmt"
	"kubecfg/config"
	"os"
)

func CreateConfigProjectSymlink() int {
	if !existingPath(configDir()) {
		fmt.Printf("project config directory '%s' is not existing\n", configDir())
		return 1
	}

	return createSymlink(unselectedProjectFile(), configSymlink())
}

func UpdateConfigProjectSymlinkToUnselected() int {
	if !existingPath(configDir()) {
		fmt.Printf("project config directory '%s' is not existing\n", configDir())
		return 1
	}

	os.Remove(configSymlink())
	return createSymlink(unselectedProjectFile(), configSymlink())
}

func UpdateConfigProjectSymlink(projectName string) int {
	os.Remove(configSymlink())
	return createSymlink(projectSymlink(projectName), configSymlink())
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

func UpdateProjectContextSymlink(projectName string, contextName config.Context) int {
	os.Remove(projectSymlink(projectName))
	return createSymlink(projectContextFile(projectName, string(contextName)), projectSymlink(projectName))
}

func UpdateProjectContextSymlinkToUnselected(projectName string) int {
	if !existingPath(projectDir(projectName)) {
		fmt.Printf("project config directory '%s' is not existing\n", projectDir(projectName))
		return 1
	}

	os.Remove(projectSymlink(projectName))
	return createSymlink(unselectedContextFile(projectName), projectSymlink(projectName))
}
