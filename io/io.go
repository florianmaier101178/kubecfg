package io

import (
	"fmt"
	"log"
	"os"
)

func IllegalConfigurationSetup() bool {
	return !existingPath(configDir()) || !existingPath(configFile())
}

func configDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.kubecfg", home)
}

func configFile() string {
	return fmt.Sprintf("%s/config.json", configDir())
}

func unselectedProjectFile() string {
	return fmt.Sprintf("%s/unselected", configDir())
}

func configSymlink() string {
	return fmt.Sprintf("%s/config", configDir())
}

func projectDir(projectName string) string {
	return fmt.Sprintf("%s/%s", configDir(), projectName)
}

func existingPath(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}

func unselectedContextFile(projectName string) string {
	return fmt.Sprintf("%s/%s/unselected", configDir(), projectName)
}

func projectSymlink(projectName string) string {
	return fmt.Sprintf("%s/%s/config_%s", configDir(), projectName, projectName)
}
