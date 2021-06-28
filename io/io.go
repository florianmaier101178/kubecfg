package io

import (
	"fmt"
	"log"
	"os"
)

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

func existingDir(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
