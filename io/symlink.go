package io

import (
	"fmt"
	"os"
)

func CreateProjectSymlink() int {
	projectDir := fmt.Sprintf("%s", configDir())
	if !existingDir(projectDir) {
		fmt.Printf("project config directory '%s' is not existing\n", projectDir)
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
