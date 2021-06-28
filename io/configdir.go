package io

import (
	"fmt"
	"os"
)

func SetupConfigDir() int {
	if existingPath(configDir()) {
		fmt.Println("config directory '~/.kubecfg' is existing and will be used")
		return 1
	}
	if _, err := os.Stat(configDir()); os.IsNotExist(err) {
		os.Mkdir(configDir(), 0755)
	}
	return 0
}

func CreateProjectDir(projectName string) int {
	projectDir := projectDir(projectName)
	if existingPath(projectDir) {
		fmt.Printf("project directory '%s' is already existing\n", projectDir)
		return 1
	}
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		os.Mkdir(projectDir, 0755)
	}
	return 0
}
